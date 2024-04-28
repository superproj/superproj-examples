package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认的 Gin 引擎
	router := gin.Default()

	// 设置组员组 "/api"
	apiGroup := router.Group("/api")
	{
		// 在 "/api" 组员组中设置路由
		apiGroup.GET("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Get users",
			})
		})
		apiGroup.POST("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Create user",
			})
		})
	}

	// 设置根路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, welcome to the root path",
		})
	})

	// 启动 Gin 服务器
	router.Run(":8080")
}
