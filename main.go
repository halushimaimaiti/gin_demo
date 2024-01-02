package main

import (
	"gin_test/config"
	"gin_test/middleware"
	"gin_test/models"
	"gin_test/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	// 使用 CORS 中间件
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}                                       // 允许所有来源
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的请求方法
	router.Use(cors.New(corsConfig))
	middleware.InitMiddleWare(router)
	routers.SetupRoutes(router)
	config.InitMysql()
	config.InitRedis()
	config.InitLog()
	models.InitDataBase()
	router.Run(":8000")
}
