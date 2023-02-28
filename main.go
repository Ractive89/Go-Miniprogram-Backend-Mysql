package main

import (
	"dayang/api"
	"dayang/conf"
	"dayang/middleware"
	"dayang/services"
	"dayang/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载启动数据库
	dayangDB := conf.Init()

	// 加载Logger
	sugarLogger := utils.InitLogger()
	defer sugarLogger.Sync()

	// 控制器
	customerService := services.NewCustomerService(dayangDB)
	customerController := api.NewCustomerController(customerService, sugarLogger)

	// 创建Gin
	//gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(middleware.TlsHandler())
	server.SetTrustedProxies(nil)
	customerController.RegisterRoutes(server)

	// 运行
	log.Fatal(server.RunTLS(conf.HttpPort, conf.SSLPem, conf.SSLKey))

}
