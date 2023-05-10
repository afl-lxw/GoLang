package main

import (
	"Golang/database"
	"Golang/routes"
)

func main() {
	//tests.InitTest()
	// 初始化数据库连接
	database.InitDB()
	defer database.CloseDB()

	routes.InitRouter()
}
