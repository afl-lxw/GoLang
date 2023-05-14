package login

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
	"net/http"
)

type LoginType struct {
	db *gorm.DB
}

func NewLogin(db *gorm.DB) *LoginType {
	return &LoginType{db: db}
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
	if !CheckCaptcha(mobile, captcha) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "验证码错误",
		})
		return
	}
	// TODO: 验证用户名和密码是否正确

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"message": "login successfully"})
}

func (h *LoginType) LoginOut(c *gin.Context) {
	c.JSON(200, gin.H{"message": "退出登录"})

	return
}

func (h *LoginType) Captcha(c *gin.Context) {
	//var store base64Captcha.Store = redis.RedisStore{}
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	if id, b64s, err := captcha.Generate(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"captcha_id": id,
			"image_data": b64s,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// CheckCaptcha 验证验证码是否正确
func CheckCaptcha(mobile, captcha string) bool {
	// 从缓存或数据库中查找验证码
	//storedCaptcha := GetCaptchaFromCache(mobile)

	// 比对验证码是否一致
	//return storedCaptcha == captcha
	return true
}
