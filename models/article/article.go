package article

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title     string    `gorm:"column:title; comment:'标题'; type:varchar(100); not null"`
	Content   string    `gorm:"column:content; comment:'内容'; type:longtext; not null"`
	Author    string    `gorm:"type:varchar(20); not null"`
	AuthorID  uint      `gorm:"type:int(11); unsigned; not null"`
	CreatedAt time.Time `gorm:"column:create_at; comment:'创建时间'; type:varchar(20); not null"`
	UpdatedAt time.Time `gorm:"column:update_at; comment:'更改时间'; type:varchar(20); not null"`
	Heart     int       `gorm:"column:heart; comment:'点赞数'; type:int(11); not null"`
	Browse    int       `gorm:"column:browse; comment:'浏览数'; type:int(11); not null"`
	IsDelete  int       `gorm:"column:is_delete; comment:'是否删除'; type:int(11); not null"`
	Address   string    `gorm:"column:address; comment:'地址'; type:varchar(100); not null"`
}
