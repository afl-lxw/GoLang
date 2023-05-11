package user

import "time"

type Usertype struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Age      int       `json:"age"`
	Gender   int       `json:"gender"`
	Mobile   string    `json:"mobile"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	UserId   int       `json:"user_id"`
	Password string    `json:"password"`
	IsDelete int       `json:"is_delete"`
	Salt     string    `json:"salt"`
}
