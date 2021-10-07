package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("0.0.0.0:5555") // 监听并在 0.0.0.0:8080 上启动服务
}
