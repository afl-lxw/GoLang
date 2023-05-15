package login

import (
	"Golang/config"
	"Golang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type LoginType struct {
	db     *gorm.DB
	config *config.Configure
}

func NewLogin(db *gorm.DB) *LoginType {
	redisClient := &config.Configure{
		Redis:       &config.RedisConfig{},
		RedisClient: &config.Redis{},
	}
	return &LoginType{db: db, config: redisClient}
}

type LoginForm struct {
	Mobile   string `form:"mobile" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}

func (h *LoginType) Login(c *gin.Context) {
	mobile := c.PostForm("mobile")
	//password := c.PostForm("password")
	captcha := c.PostForm("captcha")
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 验证验证码是否正确
	// 验证验证码是否正确
	fmt.Println("captcha: ", captcha)
	fmt.Println("mobile: ", mobile)

	// TODO: 验证用户名和密码是否正确

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"message": "login successfully"})
}

func (h *LoginType) LoginOut(c *gin.Context) {
	c.JSON(200, gin.H{"message": "退出登录"})

	return
}

func (h *LoginType) Captcha(c *gin.Context) {

	if id, b64s, err := utils.MakeCaptcha(); err == nil {
		err := h.config.RedisClient.Client.Set("captcha_id", id, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "验证码生成失败",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"captcha_id": id,
			"image_data": b64s,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
