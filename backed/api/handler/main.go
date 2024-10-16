package handler

import (
	"backed/app/core/consts"
	"backed/app/core/variable"
	"backed/app/models"
	"backed/app/utils"
	LogisticsManagement "backed/internal/contract"
	"backed/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"net/http"
	"strconv"
)

// CreatePackage 创建快递
func CreatePackage(ctx *gin.Context) {
	// 获取请求参数
	var requestData models.LogisticsInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}
	// 将包裹状态设置为待付款
	requestData.Status = 0
	// 将包裹ID设置为uuid
	requestData.PackageId = uuid.New().String()
	// 创建支付
	var tradeNo = fmt.Sprintf("%d", xid.Next())
	requestData.TradeNo = tradeNo

	// 创建数据库对象
	packageDao := database.LogisticsInfo
	// 创建包裹
	err = packageDao.WithContext(ctx).Create(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "创建包裹失败: " + err.Error(),
		})
		return
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = variable.Viper.GetString("alipay.k_server_domain") + "/third_party/notify"
	p.ReturnURL = variable.Viper.GetString("alipay.k_server_domain") + "/third_party/callback"
	p.Subject = "区块链物流平台订单:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = consts.PackagePrice
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	p.Body = requestData.PackageId

	url, err := variable.AliClient.TradePagePay(p)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "生成支付链接失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "包裹订单创建成功",
		"url":     url.String(),
	})
}

// GetPackageOrderByUser 获取用户订单
func GetPackageOrderByUser(ctx *gin.Context) {
	pageString := ctx.Query("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	sizeString := ctx.Query("size")
	size, err := strconv.ParseInt(sizeString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 从用户上下文中获取用户token
	email, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	// 首先判断用户身份
	userDao := database.UserInfo
	first, err := userDao.WithContext(ctx).
		Where(userDao.Email.Eq(email)).
		First()

	// 如果用户身份是用户 则将用户所有订单信息返回
	if first.Identity == 0 {
		// 创建数据库对象
		packageDao := database.LogisticsInfo
		// 查询订单
		packages, err := packageDao.WithContext(ctx).
			Where(packageDao.Email.Eq(email)).
			Offset((int(page) - 1) * int(size)).
			Limit(int(size)).
			Find()

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "查询订单失败: " + err.Error(),
			})
			return
		}

		if len(packages) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "暂时还没有数据",
			})
			return
		}

		var packageInfos []LogisticsManagement.LogisticsManagementPackage
		hashDao := database.TradeHash

		for _, element := range packages {
			s, err := variable.ContractSession.Packages(element.PackageId)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code":    500,
					"message": "获取包裹信息失败: " + err.Error(),
				})
				return
			}

			first, _ := hashDao.WithContext(ctx).Order(hashDao.CreatedAt.Desc()).Where(hashDao.PackageId.Eq(s.PackageId)).First()

			if first == nil {
				first = &models.TradeHash{
					TransactionHash: "",
				}
			}

			packageInfos = append(packageInfos, LogisticsManagement.LogisticsManagementPackage{
				PackageId:        s.PackageId,
				Sender:           s.Sender,
				Receiver:         s.Receiver,
				SenderAddress:    s.SenderAddress,
				ReceiverAddress:  s.ReceiverAddress,
				CurrentAddress:   first.TransactionHash,
				SupplyChainAdmin: s.SupplyChainAdmin,
				Status:           s.Status,
			})

		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查询订单成功",
			"data":    packageInfos,
		})
	} else {
		// 获取物流管理员的所有订单
		packages, err := variable.ContractSession.GetPackagesByAdmin(email)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取物流管理员订单信息: " + err.Error(),
			})
			return
		}

		if len(packages) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "暂时没有数据",
			})
			return
		}

		startIndex, endIndex := utils.CalculatePaginationIndexes(len(packages), int(page), int(size))

		hashDao := database.TradeHash
		var packageInfos []LogisticsManagement.LogisticsManagementPackage
		for i := startIndex; i >= endIndex; i-- {
			first, _ := hashDao.WithContext(ctx).Order(hashDao.CreatedAt.Desc()).Where(hashDao.PackageId.Eq(packages[i].PackageId)).First()

			if first == nil {
				first = &models.TradeHash{
					TransactionHash: "",
				}
			}
			packages[i].CurrentAddress = first.TransactionHash
			packageInfos = append(packageInfos, packages[i])
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取包裹信息成功",
			"data":    packageInfos,
		})
	}

}

// CancelPackageOrder 取消订单
func CancelPackageOrder(ctx *gin.Context) {
	// 获取请求参数
	var requestData models.LogisticsInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	// 从用户上下文中获取用户token
	email, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	packageDao := database.LogisticsInfo
	// 查询订单
	packageInfo, err := packageDao.WithContext(ctx).Where(packageDao.PackageId.Eq(requestData.PackageId)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "查询订单信息失败: " + err.Error(),
		})
		return
	}

	if packageInfo.Email != email {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "非物流发送者，不能取消物流",
		})
		return
	}

	if packageInfo.Status != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "当前订单状态不能取消",
		})
		return
	}

	_, _, err = variable.ContractSession.CancelAwaitPackage(requestData.PackageId)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "取消订单失败: " + err.Error(),
		})
		return
	}

	// 更新物流状态
	packageInfo.Status = 4
	err = packageDao.WithContext(ctx).Save(packageInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新物流状态失败: " + err.Error(),
		})
		return
	}

	variable.RefundChannel <- packageInfo.TradeNo

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "取消订单成功",
	})

}

// GetPackageOrderInTransit 获取用户所有状态在运输中的物流
func GetPackageOrderInTransit(ctx *gin.Context) {
	// 获取用户用户名
	email, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户用户名失败：" + err.Error(),
		})
		return
	}

	// 获取用户所有在运输中的物流
	packageDao := database.LogisticsInfo
	packages, err := packageDao.WithContext(ctx).
		Where(packageDao.Status.Eq(1), packageDao.Email.Eq(email)).
		Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取包裹信息失败：" + err.Error(),
		})
		return
	}

	if len(packages) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "暂时没有数据",
		})
		return
	}

	var packageInfos []LogisticsManagement.LogisticsManagementPackage
	for _, element := range packages {
		s, err := variable.ContractSession.Packages(element.PackageId)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取包裹信息失败: " + err.Error(),
			})
			return
		}
		packageInfos = append(packageInfos, LogisticsManagement.LogisticsManagementPackage{
			PackageId:        s.PackageId,
			Sender:           s.Sender,
			Receiver:         s.Receiver,
			SenderAddress:    s.SenderAddress,
			ReceiverAddress:  s.ReceiverAddress,
			CurrentAddress:   s.CurrentAddress,
			SupplyChainAdmin: s.SupplyChainAdmin,
			Status:           s.Status,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取包裹信息成功",
		"data":    packageInfos,
	})

}

// GetTracePackage 查询物流信息
func GetTracePackage(ctx *gin.Context) {
	// 获取物流ID
	packageId := ctx.Query("package_id")
	// 获取物流详情
	packages, err := variable.ContractSession.Packages(packageId)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取物流信息失败: " + err.Error(),
		})
		return
	}

	// 获取物流的出发点坐标和当前位置坐标点和终点位置坐标点
	startPoint, err := utils.SelectMapInfo(variable.Viper.GetString("mapApiKey"), packages.SenderAddress)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取出发物流坐标失败: " + err.Error(),
		})
		return
	}

	currentPoint, err := utils.SelectMapInfo(variable.Viper.GetString("mapApiKey"), packages.CurrentAddress)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取当前物流坐标失败: " + err.Error(),
		})
		return
	}

	endPoint, err := utils.SelectMapInfo(variable.Viper.GetString("mapApiKey"), packages.ReceiverAddress)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取接收物流坐标失败: " + err.Error(),
		})
		return
	}

	// 将上面这些坐标转成切片float
	startPoints, err := utils.TransferStringToFloat(startPoint)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "转换出发物流坐标失败: " + err.Error(),
		})
		return
	}
	currentPoints, err := utils.TransferStringToFloat(currentPoint)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "转换当前物流坐标失败: " + err.Error(),
		})
		return
	}
	endPoints, err := utils.TransferStringToFloat(endPoint)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "转换接收物流坐标失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取物流坐标成功",
		"start":   startPoints,
		"current": currentPoints,
		"end":     endPoints,
	})
}

// GetIndexInfo 获取首页信息
func GetIndexInfos(ctx *gin.Context) {
	// 需要获取近期两笔交易
	packageDao := database.LogisticsInfo
	packages, err := packageDao.WithContext(ctx).Limit(2).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取物流信息失败: " + err.Error(),
		})
		return
	}

	var packageInfos []LogisticsManagement.LogisticsManagementPackage
	hashDao := database.TradeHash
	for _, element := range packages {
		s, err := variable.ContractSession.Packages(element.PackageId)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取包裹信息失败: " + err.Error(),
			})
			return
		}

		first, _ := hashDao.WithContext(ctx).Order(hashDao.CreatedAt.Desc()).Where(hashDao.PackageId.Eq(s.PackageId)).First()

		if first == nil {
			first = &models.TradeHash{
				TransactionHash: "",
			}
		}
		packageInfos = append(packageInfos, LogisticsManagement.LogisticsManagementPackage{
			PackageId:        s.PackageId,
			Sender:           s.Sender,
			Receiver:         s.Receiver,
			SenderAddress:    s.SenderAddress,
			ReceiverAddress:  s.ReceiverAddress,
			CurrentAddress:   s.CurrentAddress,
			SupplyChainAdmin: first.TransactionHash,
			Status:           s.Status,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取物流信息成功",
		"data":    packageInfos,
	})
}
