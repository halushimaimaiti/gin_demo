package controller

import (
	"fmt"
	"gin_test/config"
	"gin_test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	name := c.Query("name")
	fmt.Println(name)
	fmt.Println(len(name))
	people := service.GetUser()
	c.JSON(http.StatusOK, people)
	config.Log.Error("hello")
}

func LoginController(c *gin.Context) {
	fmt.Println(121)
	username := c.PostForm("username")
	fmt.Println(username)
	c.JSON(http.StatusOK, nil)
	config.Log.Error("hello")
}
