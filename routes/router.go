package routes

import (
	"Golang/database"
	_ "Golang/docs" // 导入 API 文档
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Engine      *gin.Engine
	RouterGroup *gin.RouterGroup
}

func InitRouter() *Router {
	database.InitDB()
	//defer database.CloseDB()
	// 创建一个路由
	//engine := gin.Default()
	engine := gin.New()
	//r := New()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 初始化 RouterGroup
	routerGroup := engine.Group("api/v1")

	// 注册 v1 版本的路由处理函数
	RegisterV1HandlersUser(routerGroup)

	//engine.Run("127.0.0.1:8080")

	// 返回 Router 实例
	return &Router{
		Engine:      engine,
		RouterGroup: routerGroup,
	}
}
