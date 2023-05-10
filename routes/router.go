package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 创建一个路由
	r := gin.Default()
	//r := New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	// 路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			authFlag, err := c.Get("auth")
			if !err {
				authFlag = false
			}
			fmt.Println(authFlag)
			c.JSON(200, gin.H{
				"message": "v1 login",
			})
		})
	}

	r.Run("127.0.0.1:8080")
}
