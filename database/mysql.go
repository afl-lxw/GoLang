package database

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

// InitDB 初始化数据库连接
func InitDB() {
	// 使用 flag 包从命令行中读取配置文件路径，默认为 config.toml
	var configPath string
	flag.StringVar(&configPath, "conf", "./conf/config.yaml", "path to config file")
	flag.Parse()

	// 使用 viper 读取配置文件
	viper.SetConfigFile(configPath)
	configErr := viper.ReadInConfig()
	if configErr != nil {
		// 处理配置文件读取错误
		log.Fatal(configErr)
	}

	// 获取配置项
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	name := viper.GetString("database.name")

	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, name)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	db.Close()
}
