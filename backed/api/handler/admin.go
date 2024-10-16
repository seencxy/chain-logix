package handler

import (
	"backed/app/core/variable"
	"backed/app/models"
	"backed/app/utils"
	LogisticsManagement "backed/internal/contract"
	"backed/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllUser 获取所有用户
func GetAllUser(ctx *gin.Context) {
	// 获取请求参数
	var requestData models.GetAllUserRequest
	var err error
	requestData.Page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
	requestData.Number, err = strconv.ParseInt(ctx.Query("number"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取所有用户信息
	userDao := database.UserInfo
	find, err := userDao.WithContext(ctx).
		Offset(int(requestData.Number) * (int(requestData.Page) - 1)).
		Limit(int(requestData.Number)).
		Find()

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取所有用户信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户信息成功",
		"data":    find,
	})
}

// UpdateAccountStatus 修改用户状态
func UpdateAccountStatus(ctx *gin.Context) {
	// 实例化请求参数
	var requestData models.UserInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
	}

	email := requestData.Email

	if len(email) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + "email不能为空",
		})
		return
	}

	// 更新用户账户
	userDao := database.UserInfo
	_, err = userDao.WithContext(ctx).
		Where(userDao.Email.Eq(email)).
		Update(userDao.Status, requestData.Status)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "将用户账号停用失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "将用户设置为停用成功",
	})
}

// GetPackageInfo 获取包裹信息
func GetPackageInfo(ctx *gin.Context) {
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
	// 获取所有订单信息
	packages, err := variable.ContractSession.GetAllPackages()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取所有包裹信息失败: " + err.Error(),
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

	// 对订单信息进行分页
	startIndex, endIndex := utils.CalculatePaginationIndexes(len(packages), int(page), int(size))

	var packageInfos []LogisticsManagement.LogisticsManagementPackage
	for i := startIndex; i >= endIndex; i-- {
		packageInfos = append(packageInfos, packages[i])
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取包裹信息成功",
		"data":    packageInfos,
	})
}

// ReceivePackage 物流管理员接收物流
func ReceivePackage(ctx *gin.Context) {
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

	// 确定用户身份是物流管理员
	query := database.UserInfo
	userInfo, err := query.WithContext(ctx).Where(query.Email.Eq(requestData.Email)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户非物流管理员！",
		})
		return
	}

	_, transaction, err := variable.ContractSession.SupplyAdminAcceptPackage(requestData.Email, requestData.PackageId)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定物流管理员失败: " + err.Error(),
		})
		return
	}

	// 为此包裹ID添加最新hash
	hashDao := database.TradeHash
	var hashModel models.TradeHash
	hashModel.PackageId = requestData.PackageId
	hashModel.TransactionHash = transaction.TransactionHash
	err = hashDao.WithContext(ctx).Create(&hashModel)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建最新交易hash: " + err.Error(),
		})
		return
	}

	// 更新包裹状态
	packageDao := database.LogisticsInfo
	_, err = packageDao.WithContext(ctx).
		Where(packageDao.PackageId.Eq(requestData.PackageId)).
		Update(packageDao.Status, 2)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新包裹状态失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新包裹状态成功",
	})
}

// UpdatePackageStatus 物流管理员更新物流状况
func UpdatePackageStatus(ctx *gin.Context) {
	// 获取请求参数
	var requestData models.UpdatePackageStatusRequest
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	// 确定用户身份是物流管理员
	query := database.UserInfo
	userInfo, err := query.WithContext(ctx).
		Where(query.Email.Eq(requestData.Email)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取请求参数失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户非物流管理员！",
		})
		return
	}

	// 更新包裹状态
	_, transaction, err := variable.ContractSession.
		UpdateCurrentAddress(userInfo.Email, requestData.PackageId, requestData.CurrentAddress)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新物流状态失败: " + err.Error(),
		})
		return
	}

	// 为此包裹ID添加最新hash
	hashDao := database.TradeHash
	var hashModel models.TradeHash
	hashModel.PackageId = requestData.PackageId
	hashModel.TransactionHash = transaction.TransactionHash
	err = hashDao.WithContext(ctx).Create(&hashModel)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建最新交易hash: " + err.Error(),
		})
		return
	}

	// 更新数据中物流状态
	packageDao := database.LogisticsInfo
	// 首先获取物流信息
	packageInfo, err := packageDao.WithContext(ctx).
		Where(packageDao.PackageId.Eq(requestData.PackageId)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取物流信息失败: " + err.Error(),
		})
		return
	}

	var status int64
	// 代表物流已经送达了
	if packageInfo.ReceiveAddress == requestData.CurrentAddress {
		status = 2
	} else {
		status = 1
	}

	_, err = packageDao.WithContext(ctx).
		Where(packageDao.PackageId.Eq(requestData.PackageId)).
		Update(packageDao.Status, status)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新物流状态失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新物流状态成功",
	})
}
