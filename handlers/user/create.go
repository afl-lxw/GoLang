package user

import (
	userOrm "Golang/models/user"
	"Golang/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
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

	if err := c.ShouldBind(h.userType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "请求错误"})
		return
	}
	validate := validator.New()
	validate.RegisterValidation("complexity", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		// 定义密码复杂度要求
		minLength := 8 // 最小长度为8
		// 验证密码长度
		if len(password) < minLength {
			return false
		}
		// 验证是否包含大写字母
		if !containsUppercase(password) {
			return false
		}
		// 验证是否包含小写字母
		if !containsLowercase(password) {
			return false
		}
		// 验证是否包含数字
		if !containsDigit(password) {
			return false
		}
		// 验证是否包含特殊字符
		if !containsSpecialCharacters(password) {
			return false
		}
		return true
	})
	err := validate.Struct(h.userType)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "请求错误"})
			return
		}
	} else {
		fmt.Println("Validation passed")
	}
	h.userType.UserId = uuid.New()
	// 校验并设置性别
	if h.userType.Sex == "" {
		h.userType.Sex = "其他"
	}
	if h.userType.Sex != "男" && h.userType.Sex != "女" && h.userType.Sex != "其他" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sex"})
		return
	}

	if utils.ValidatePhone(strconv.FormatInt(h.userType.Mobile, 10)) {
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "手机号格式错误"})
		return
	}

	h.userType.CreatedAt = time.Now()
	h.userType.UpdatedAt = time.Now()

	// 生成一个长度为 16 的随机盐

	salt, err := utils.RandSalt()
	if err != nil {
		fmt.Printf("生成随机盐失败: %v", err)
	}
	h.userType.Salt = salt
	password, errPass := utils.PasswordHash(h.userType.Password)

	if errPass != nil {
		fmt.Printf("密码加密失败: %v", errPass)
	}
	h.userType.Password = password

	result := h.db.Create(h.userType)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error, "message": "数据创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": true, "Message": "创建成功"})
}

// 判断字符串是否包含大写字母
func containsUppercase(s string) bool {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

// 判断字符串是否包含小写字母
func containsLowercase(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			return true
		}
	}
	return false
}

// 判断字符串是否包含数字
func containsDigit(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

// 判断字符串是否包含特殊字符
func containsSpecialCharacters(s string) bool {
	specialCharacters := []rune{'!', '@', '#', '$', '%', '^', '&', '*'}
	for _, char := range s {
		for _, specialChar := range specialCharacters {
			if char == specialChar {
				return true
			}
		}
	}
	return false
}
