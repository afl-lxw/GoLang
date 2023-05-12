package user

import "time"

type User struct {
	Id       int64     `gorm:"column:id"`
	Username string    `gorm:"column:username"`
	Age      int       `gorm:"column:age"`
	Gender   string    `gorm:"column:gender"`
	Mobile   string    `gorm:"column:mobile"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
	UserId   string    `gorm:"column:user_id"`
	Password string    `gorm:"column:password,omitempty "`
	IsDelete int       `gorm:"column:isDelete"`
	Salt     string    `gorm:"column:salt,omitempty "`
}
