package routes

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	RouterGroup *gin.RouterGroup
}

func InitRouter() *Router {
	// 创建一个路由
	//engine := gin.Default()
	engine := gin.New()
	//r := New()

	// 初始化 RouterGroup
	routerGroup := engine.Group("v1")

	// 注册 v1 版本的路由处理函数
	RegisterV1HandlersUser(routerGroup)

	//engine.Run("127.0.0.1:8080")

	// 返回 Router 实例
	return &Router{
		Engine:      engine,
		RouterGroup: routerGroup,
	}
}
