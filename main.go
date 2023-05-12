package main

import (
	"Golang/routes"
	"fmt"
)

// @title goweb project
// @version 1.0
// @description this is goweb server.
// @host 127.0.0.1:6912
// @BasePath /api/v1
func main() {
	router := routes.InitRouter()
	// 启动服务
	if err := router.Engine.Run(":8080"); err != nil {
		fmt.Println("Server Error: ", err)
	}
}
