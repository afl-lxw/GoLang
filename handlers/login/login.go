package login

import (
	"Golang/config"
	"Golang/utils"
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
	Id       string `form:"id" binding:"required"`
}

func (h *LoginType) Login(c *gin.Context) {
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	captcha := c.PostForm("captcha")
	id := c.PostForm("id")
	println("id: ", id, "captcha: ", captcha, "mobile: ", mobile, "password: ", password)
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 验证验证码是否正确
	// 验证验证码是否正确
	capValue := utils.VerifyCaptcha(form.Id, form.Captcha)
	println("capValue: ", capValue)
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
		//err := h.config.RedisClient.Client.Set("captcha_id", id, 0)
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
