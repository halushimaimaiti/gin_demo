package routers

import (
	"gin_test/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", controller.HomeController)
	router.POST("/login", controller.LoginController)
}
