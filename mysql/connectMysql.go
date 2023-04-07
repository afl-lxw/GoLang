package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 用户名：密码^@tcp(地址:3306)/数据库
// 表结构
type info struct {
	id     int    `db:"id"`
	name   string `db:"name"`
	author string `db:"author"`
}

func Mysql_start() {
	db, _ := sql.Open("mysql", "root:liuxiaowen@(127.0.0.1:3306)/node?charset=utf8&parseTime=true")
	defer db.Close() //关闭数据库
	err := db.Ping() //连接数据库
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	//查询表
	rows, errs := db.Query("select * from tags")
	if errs != nil {
		fmt.Println("数据库连接失败------", errs)
		return
	}
	fmt.Println("\n", rows)
	//遍历打印
	for rows.Next() {
		var id int
		var tagname string
		rows.Scan(&id, &tagname)
		fmt.Println(id, tagname)
	}
	//用完关闭
	rows.Close()
}
