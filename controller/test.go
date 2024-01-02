package controller

import (
	"gin_test/config"
	"gin_test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {

	people := service.GetUser()
	c.JSON(http.StatusOK, people)
	config.Log.Error("hello")
}
