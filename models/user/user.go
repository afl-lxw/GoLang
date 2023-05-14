package user

import (
	"github.com/google/uuid"
	"time"
)

type Gender string

const (
	Male   Gender = "男"
	Female Gender = "女"
	other  Gender = "其他"
)

type User struct {
	//ID uint `gorm:"column:id;type:int(11) unsigned auto_increment;not null;primaryKey;comment:记录ID"`
	//Id       int64  `json:"id"  gorm:"primaryKey; autoIncrement; "`
	Username  string    `json:"username" binding:"required" form:"username" gorm:"column:username;type:varchar(200); comment:'用户名';uniqueIndex"`
	Age       int       `json:"age" form:"age" gorm:"column:age; comment:'年龄'; default:18 "`
	Sex       Gender    `json:"sex" form:"sex"  gorm:"column:sex; comment: '性别'"`
	Mobile    int64     `json:"mobile" binding:"required" form:"mobile"  gorm:"column:mobile; comment: '电话';uniqueIndex"`
	Email     string    `json:"email" binding:"required" form:"email"  gorm:"column:email;comment:'用户邮箱'"`
	Password  string    `json:"password" binding:"required" form:"password"  gorm:"column:password;omitempty, comment:'密码'"`
	CreatedAt time.Time `json:"created_at"    gorm:"column:create_at; comment:'创建时间'"`
	UpdatedAt time.Time `json:"updated_at"  gorm:"column:update_at; comment:'修改时间'"`
	UserId    uuid.UUID `json:"user_id"  gorm:"column:user_id; comment: '用户Uid'"`
	IsDelete  int       `json:"is_delete"   gorm:"column:is_delete; comment:'软删除'"`
	Salt      []byte    `json:"salt"  gorm:"column:salt; comment:'密码盐' "`
}

type GormUser struct {
	ID uint `gorm:"column:id;type:int(11) unsigned auto_increment;not null;primaryKey;comment:记录ID"`
	*User
}

func (GormUser) TableName() string {
	return "users"
}
