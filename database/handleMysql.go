package database

import (
	articleOrm "Golang/models/article"
	userOrm "Golang/models/user"
	"gorm.io/gorm"
)

func CheckDatabase(db *gorm.DB) {
	// 执行迁移操作
	if err := db.AutoMigrate(&userOrm.GormUser{}); err != nil {
		panic("failed to migrate database")
	}
	if err := db.AutoMigrate(&articleOrm.Article{}); err != nil {
		panic("failed to migrate database")
	}
}
