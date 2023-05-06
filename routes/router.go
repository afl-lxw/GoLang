package routes

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 用户模块
	v1User := router.Group("/v1/user")
	{
		v1User.GET("/", getUserList)
		v1User.GET("/:id", getUser)
		v1User.POST("/", createUser)
		v1User.PUT("/:id", updateUser)
		v1User.DELETE("/:id", deleteUser)
	}

	// 主页模块
	v1Home := router.Group("/v1/home")
	{
		v1Home.GET("/", getHomePage)
		v1Home.GET("/:id", getArticle)
		v1Home.POST("/", createArticle)
		v1Home.PUT("/:id", updateArticle)
		v1Home.DELETE("/:id", deleteArticle)
	}

	// 登录模块
	v1Login := router.Group("/v1/login")
	{
		v1Login.POST("/", login)
		v1Login.POST("/logout", logout)
	}

	return router
}

func getUserList(c *gin.Context) {
	// TODO: 获取用户列表
}

func getUser(c *gin.Context) {
	// TODO: 获取用户信息
}

func createUser(c *gin.Context) {
	// TODO: 创建用户
}

func updateUser(c *gin.Context) {
	// TODO: 更新用户信息
}

func deleteUser(c *gin.Context) {
	// TODO: 删除用户
}

func getHomePage(c *gin.Context) {
	// TODO: 获取主页信息
}

func getArticle(c *gin.Context) {
	// TODO: 获取文章信息
}

func createArticle(c *gin.Context) {
	// TODO: 创建文章
}

func updateArticle(c *gin.Context) {
	// TODO: 更新文章信息
}

func deleteArticle(c *gin.Context) {
	// TODO: 删除文章
}

func login(c *gin.Context) {
	// TODO: 登录
}

func logout(c *gin.Context) {
	// TODO: 登出
}
