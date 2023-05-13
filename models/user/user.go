package user

import "time"

type Gender string

const (
	Male   Gender = "男"
	Female Gender = "女"
	other  Gender = "其他"
)

type User struct {
	*UserForm
	CreateAt time.Time `json:"create_at"    gorm:"column:create_at; comment:'创建时间'"`
	UpdateAt time.Time `json:"update_at"  gorm:"column:update_at; comment:'修改时间'"`
	UserId   string    `json:"user_id"  gorm:"column:user_id; comment: '用户Uid'"`
	IsDelete int       `json:"is_delete"   gorm:"column:is_delete; comment:'软删除'"`
	Salt     string    `json:"salt"  gorm:"column:salt; comment:'密码盐' "`
}

type UserForm struct {
	Id       int64  `json:"id" form:"id" gorm:"primaryKey;autoIncrement; column:id; comment:'用户id';uniqueIndex"`
	Username string `json:"username" binding:"required" form:"username" gorm:"column:username;type:varchar(200); comment:'用户名'"`
	Age      int    `json:"age" form:"age" gorm:"column:age; comment:'年龄'"`
	Sex      Gender `json:"sex" form:"sex"  gorm:"column:sex;type:enum('男', '女', '其他'); default:'other'; comment: '性别'"`
	Mobile   string `json:"mobile" binding:"required" form:"mobile"  gorm:"column:mobile; comment: '电话'"`
	Email    string `json:"email" binding:"required" form:"email"  gorm:"column:email;comment:'用户邮箱'"`
	Password string `json:"password" binding:"required" form:"password"  gorm:"column:password;omitempty, comment:'密码'"`
}
