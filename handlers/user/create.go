package user

import (
	userOrm "Golang/models/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserCreate struct {
	db       *gorm.DB
	userType *userOrm.User
}

func CreateUser(db *gorm.DB) *UserCreate {
	u := &userOrm.User{}
	return &UserCreate{db: db, userType: u}
}

// UserCreate @Summary 创建用户
// @Description 创建一个新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param username body string true "用户名"
// @Param mobile body string true "手机号码"
// @Param password body string true "密码"
// @Param age body int false "年龄"
// @Param gender body int false "性别"
// @Success 200 {object} UserResponse "返回创建成功的用户信息"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /users [post]
func (h *UserCreate) UserCreate(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Sprintf("username:%s,password:%s,==========type:%s", username, password, types)

	for key, values := range c.Request.Form {
		fmt.Printf("%s:---- %v\n", key, values)
	}

	for key, values := range c.Request.PostForm {
		fmt.Printf("%s:=== %v\n", key, values)
	}

	if err := c.ShouldBindJSON(h.userType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "请求错误"})
		return
	}
	fmt.Println("username:", h.userType.Username)
	fmt.Println("password:", h.userType.Password)

	h.db.AutoMigrate(h.userType)

	h.db.Create(&h.userType)

	c.JSON(http.StatusOK, gin.H{"Data": true, "Message": "创建成功"})
}
