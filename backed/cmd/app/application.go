package main

import (
	"backed/api/handler"
	"backed/api/router"
	"backed/app/core/variable"
	_ "backed/bootstrap"
	"github.com/gin-gonic/gin"
)

// 项目主入口
func main() {
	// 引入gin引擎
	engine := gin.Default()

	// 加载路由
	router.MainRouter(engine)

	// 项目总协程
	go handler.MainGoroutine()
	// 合约事件监听
	go handler.ListenEvent()

	// 监听路由
	engine.Run(":" + variable.Viper.GetString("application_port"))
}
