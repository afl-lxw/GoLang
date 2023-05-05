package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func IndexMiddle(c *gin.Context) {
	start := time.Now()
	fmt.Printf("index middle --> %s \n", start)
	//go func(c.Copy()) {}()  只能使用c的Copy()方法
	c.Next()
	end := time.Since(start)
	fmt.Printf("index middle times --> %s \n", end)
}

func AuthMiddleware(flag bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if flag {

		} else {
			fmt.Println("auth middleware")
			c.Set("auth", true)
			c.Next()
		}

	}
}

func main() {
	//tests.InitTest()
	r := gin.Default()
	//r := New()
	r.Use(AuthMiddleware(false), IndexMiddle) // 全局中间件
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	// 路由组
	v1 := r.Group("/v1", AuthMiddleware(false))
	{
		v1.GET("/login", IndexMiddle, func(c *gin.Context) {
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
