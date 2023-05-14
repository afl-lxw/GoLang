package database

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func GetDB() *gorm.DB {
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
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	CheckDatabase(db)

	fmt.Println("Successfully connected to the database!")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close()
}
