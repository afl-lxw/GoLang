package routes

import (
	"Golang/app/middleware/jwt"
	"Golang/database"
	"Golang/handlers/user"
	"github.com/gin-gonic/gin"
)

func RegisterV1HandlersUser(routerGroup *gin.RouterGroup) {
	db := database.GetDB()
	// 注册 v1 版本的路由处理函数
	userHandler := user.GetUserList(db)
	userHandlerCreate := user.CreateUser(db)
	routerGroup.GET("/users", jwt.JWTMiddleware(), userHandler.GetUserList)
	routerGroup.POST("/users", userHandlerCreate.UserCreate)
}
