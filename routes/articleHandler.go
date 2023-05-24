package routes

import (
	"Golang/database"
	"Golang/handlers/article"
	"github.com/gin-gonic/gin"
)

func RegisterV1HandlersArticle(routerGroup *gin.RouterGroup) {
	db := database.GetDB()
	// 注册 v1 版本的路由处理函数
	ArticleHandler := article.NewArticle(db)
	routerGroup.POST("/article", ArticleHandler.CreateArticle)
	routerGroup.GET("/article", ArticleHandler.CreateArticle)
	routerGroup.PATCH("/article", ArticleHandler.CreateArticle)
	routerGroup.DELETE("/article", ArticleHandler.CreateArticle)
}
