package controller

import (
	"fmt"
	"gin_test/config"
	"gin_test/service"
	"net/http"
	"strconv"

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

func FindBookByInfo(c *gin.Context) {
	word := c.Query("word")
	page := c.Query("page")
	tempPage, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "page非整数",
		})
	}
	config.Log.Error("hello")
	bookInfo := service.FindBookByInfo(word, int32(tempPage))
	c.JSON(http.StatusOK, bookInfo)
}
