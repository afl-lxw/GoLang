package routes

import (
	"Golang/database"
	"Golang/handlers/login"
	"github.com/gin-gonic/gin"
)

func RegisterV1HandlersLogin(routerGroup *gin.RouterGroup) {
	db := database.GetDB()
	// 注册 v1 版本的路由处理函数
	loginHandler := login.NewLogin(db)
	routerGroup.POST("/login", loginHandler.Login)
	routerGroup.GET("/login-out", loginHandler.LoginOut)
	routerGroup.GET("/captcha", loginHandler.Captcha)
}
