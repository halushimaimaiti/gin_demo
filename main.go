package main

import (
	"gin_test/config"
	"gin_test/middleware"
	"gin_test/models"
	"gin_test/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	middleware.InitMiddleWare(router)
	routers.SetupRoutes(router)
	config.InitMysql()
	config.InitRedis()
	config.InitLog()
	models.InitDataBase()
	router.Run(":8000")
}
