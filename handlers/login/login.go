package login

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginType struct {
	db *gorm.DB
}

func NewLogin(db *gorm.DB) *LoginType {
	return &LoginType{db: db}
}

func (h *LoginType) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login"})
	return
}

func (h *LoginType) LoginOut(c *gin.Context) {
	c.JSON(200, gin.H{"message": "退出登录"})

	return
}
