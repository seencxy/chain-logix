package router

import (
	"backed/api/handler"
	"backed/api/middleware"
	"github.com/gin-gonic/gin"
)

// MainRouter 项目总路由
func MainRouter(engine *gin.Engine) {
	// 配置全局跨域中间键
	engine.Use(middleware.CorsHandler)

	// 创建websocket连接
	engine.GET("/web_link", handler.Websocket)

	// 注册第三方路由
	thirdPartyRouter(engine.Group("/third_party"))

	// 用户模块路由
	userRouter(engine.Group("/user"))

	// 主程序路由
	mainRouter(engine.Group("/main"))

	// 管理模块路由
	adminRouter(engine.Group("/admin"))
}

// mainRouter 主程序路由
func mainRouter(router *gin.RouterGroup) {
	// 创建包裹
	router.POST("/createPackage", middleware.JwtTokenValid, handler.CreatePackage)
	// 获取用户订单
	router.GET("/getUserPackage", middleware.JwtTokenValid, handler.GetPackageOrderByUser)
	// 取消物流订单
	router.POST("/cancelPackage", middleware.JwtTokenValid, handler.CancelPackageOrder)
	// 获取用户运输中的物流
	router.GET("/getUserTransportPackage", middleware.JwtTokenValid, handler.GetPackageOrderInTransit)
	// 获取用户物流详情
	router.GET("/getUserPackageDetail", middleware.JwtTokenValid, handler.GetTracePackage)
	// 获取用户首页信息
	router.GET("/getIndexInfos", middleware.JwtTokenValid, handler.GetIndexInfos)
}

// userRouter 用户模块路由
func userRouter(router *gin.RouterGroup) {
	// 用户注册接口
	router.POST("/register", handler.Register)
	// 发送邮箱验证码接口
	router.POST("/sendCode", handler.SendEmailCode)
	// 用户登录接口
	router.POST("/login", handler.Login)
	// 重置密码接口
	router.POST("/resetPassword", handler.ResetPassword)
	// 修改用户头像接口
	router.POST("/changeUserHead", handler.ChangeUserHead)
}

// third_party 第三方路由
func thirdPartyRouter(router *gin.RouterGroup) {
	// 支付成功反馈
	router.GET("/callback", handler.AlipayCallback)
	// 支付提示
	router.POST("/notify", middleware.JwtTokenValid, handler.AlipayNotify)
}

// admin 管理模块路由
func adminRouter(router *gin.RouterGroup) {
	// 获取所有用户信息
	router.GET("/getAllUser", middleware.JwtTokenValid, handler.GetAllUser)
	// 修改用户状态
	router.POST("/UpdateAccountStatus", middleware.JwtTokenValid, handler.UpdateAccountStatus)
	// 获取所有订单信息
	router.GET("/getAllPackage", middleware.JwtTokenValid, handler.GetPackageInfo)
	// 物流管理员接收物流
	router.POST("/receivePackage", middleware.JwtTokenValid, handler.ReceivePackage)
	// 物流管理员更新物流状态
	router.POST("/updatePackageStatus", middleware.JwtTokenValid, handler.UpdatePackageStatus)
}
