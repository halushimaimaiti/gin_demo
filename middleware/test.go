package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleWare(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		// 在处理请求前执行的操作
		// 可以进行身份验证、日志记录等操作
		// 例如，在这里打印请求路径
		// path := c.Request.URL.Path
		// if path != "/api/login" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		// 	c.Abort()
		// 	return
		// }
		// 继续处理请求
		c.Next()
	})
}
