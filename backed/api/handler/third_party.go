package handler

import (
	"backed/app/core/consts"
	"backed/app/core/variable"
	"backed/app/models"
	"backed/internal/database"
	"context"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
	"log"
	"net/http"
	"time"
)

var clients = make(map[string]*Client)

// Websocket 网站之间的websocket连接
func Websocket(ctx *gin.Context) {
	email := ctx.Query("email")
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		HandshakeTimeout: 10 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}

	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "建立websocket连接失败: " + err.Error(),
		})
		return
	}
	defer conn.Close()

	// 添加客户端到已连接客户端列表
	client := &Client{email: email, conn: conn}
	clients[email] = client

	// 保持连接打开，直到发生错误或客户端关闭连接
	for {
		// 等待客户端消息或连接关闭
		if _, _, err := conn.ReadMessage(); err != nil {
			break // 如果发生错误或连接关闭，则退出循环
		}
		// 可以在这里处理任何消息，或者仅保持连接活跃
	}

	// 连接关闭或出错时，将该用户设置为离线
	delete(clients, email)
}

// SystemSendMessage 系统发布消息到用户
func SystemSendMessage(email, message string) {

	// 首先判断用户是否在线
	if _, ok := clients[email]; !ok {
		return
	}
	// 发送消息
	if err := clients[email].conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		logrus.Errorf("发送消息失败: %s", err.Error())
		return
	}
}

// Client 客户端
type Client struct {
	email string
	// websocket连接
	conn *websocket.Conn
}

// AlipayCallback 支付宝成功回调函数
func AlipayCallback(ctx *gin.Context) {
	// 获取参数
	err := ctx.Request.ParseForm()
	if err != nil {
		logrus.Error("获取请求参数失败: ", err.Error())
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}
	var outTradeNo = ctx.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := variable.AliClient.TradeQuery(p)
	if err != nil {
		logrus.Error("验证请求失败: ", err.Error())
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}
	if rsp.IsSuccess() == false {
		logrus.Error("验证订单信息失败: ", err.Error())
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}
	logrus.Info("支付成功: ", rsp)
	// 更新订单消息
	query := database.LogisticsInfo

	// 获取需要发送包裹的信息
	first, err := query.WithContext(ctx).
		Where(query.TradeNo.Eq(rsp.OutTradeNo)).
		First()
	if err != nil {
		logrus.Error("获取发送包裹信息失败: ", err.Error())
		variable.RefundChannel <- first.TradeNo
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}

	// 发布包裹
	_, transaction, err := variable.ContractSession.RegisterPackage(first.PackageId, first.SendPerson, first.ReceivePerson, first.SendDetailedAddress, first.ReceiveAddress)
	if err != nil {
		// 创建包裹失败
		logrus.Error("创建包裹失败: ", err.Error())
		variable.RefundChannel <- first.TradeNo
		ctx.Redirect(http.StatusFound, consts.FrontAddress)
		return
	}

	// 为此包裹ID添加最新hash
	hashDao := database.TradeHash
	var hashModel models.TradeHash
	hashModel.PackageId = first.PackageId
	hashModel.TransactionHash = transaction.TransactionHash
	err = hashDao.WithContext(ctx).Create(&hashModel)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建最新交易hash: " + err.Error(),
		})
		return
	}

	message := models.MessageInfo{
		UserName: first.Email,
		Message:  "您的包裹已经发出,请注意查收",
	}

	variable.MessageChannel <- message.Marshal()

	ctx.Redirect(http.StatusFound, consts.FrontAddress)
}

// AlipayNotify 支付宝提示接口
func AlipayNotify(c *gin.Context) {
	c.Request.ParseForm()

	err := variable.AliClient.VerifySign(c.Request.Form)
	if err != nil {
		logrus.Error("异步通知验证签名发生错误", err)
		return
	}

	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := variable.AliClient.TradeQuery(p)
	if err != nil {
		logrus.Error("异步通知验证订单 %s 信息发生错误: %s \n", outTradeNo, err.Error())
		return
	}
	if rsp.IsSuccess() == false {
		logrus.Error("异步通知验证订单 %s 信息发生错误: %s-%s \n", outTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	logrus.Error("订单 %s 支付成功 \n", outTradeNo)
}

// MainGoroutine 项目总携程
func MainGoroutine() {
	for {
		select {
		case data := <-variable.MessageChannel: // 需要发送消息
			// 创建实例化消息
			var message models.MessageInfo
			// 将请求消息
			err := message.Unmarshal([]byte(data))
			if err != nil {
				logrus.Error("获取请求消息失败: ", err.Error())
				continue
			}
			// 发送消息
			SystemSendMessage(message.UserName, message.Message)
		case data := <-variable.RefundChannel: // 需要退款
			result, err := variable.AliClient.TradeRefund(alipay.TradeRefund{
				OutTradeNo:   data,
				RefundAmount: consts.PackagePrice,
				RefundReason: "用户取消包裹或者系统异常",
			})

			logrus.Error("退款信息: ", result)

			if err != nil {
				logrus.Error(data+" 退款失败: ", err.Error())
				return
			}

			// 退款成功后,将订单状态设置为已退款
			query := database.LogisticsInfo
			info, err := query.WithContext(context.TODO()).Where(query.TradeNo.Eq(data)).First()
			info.Status = 4
			// 更新状态
			err = query.WithContext(context.TODO()).Save(info)
			if err != nil {
				logrus.Error("更新状态失败", err.Error())
				return
			}
			// 发送消息给用户
			message := models.MessageInfo{
				UserName: info.SendPerson,
				Message:  "您的包裹 " + info.PackageId + " 已经退款成功",
			}

			variable.MessageChannel <- message.Marshal()

			variable.AliClient.FundTransToAccountTransfer(alipay.FundTransToAccountTransfer{})
		}
	}
}

// ListenEvent 监听合约事件的发生
func ListenEvent() {
	// 获取最新区块 高度
	blockNumber, err := variable.FiscoBcos.GetBlockNumber(context.TODO())
	if err != nil {
		logrus.Error("获取最新区块高度失败: ", err.Error())
	}
	// 获取当前区块高度
	currentBlockNumber := uint64(blockNumber)
	_, err = variable.ContractSession.Contract.WatchPackageStatusUpdated(&currentBlockNumber, func(status int, logs []types.Log) {
		for _, v := range logs {
			logrus.Info("监听到事件发生")
			log, errs := variable.ContractSession.ParsePackageStatusUpdated(v) // 解析数据
			if errs != nil {
				logrus.Println(errs)
			}

			// 获取该包裹的用户
			query := database.LogisticsInfo
			info, err := query.WithContext(context.TODO()).Where(query.PackageId.Eq(log.PackageId)).First()
			if err != nil {
				logrus.Println(errs)
			}

			_, err = query.WithContext(context.TODO()).
				Where(query.PackageId.Eq(log.PackageId)).
				Update(query.Status, 1)
			if err != nil {
				logrus.Error("修改订单信息失败: ", err.Error())
			}
			// 创建消息
			message := fmt.Sprintf("包裹ID: %s, 包裹状态: %s", log.PackageId, log.NewStatus)
			newMessage := models.MessageInfo{
				Message:  message,
				UserName: info.Email,
			}

			variable.MessageChannel <- newMessage.Marshal()
		}
	})
	if err != nil {
		log.Println("监听合约失败: " + err.Error())
		return
	}
}
