package user

import (
	userOrm "Golang/models/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
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

	h.userType.UserId = uuid.New()
	// 校验并设置性别
	if h.userType.Sex == "" {
		h.userType.Sex = "其他"
	}
	if h.userType.Sex != "男" && h.userType.Sex != "女" && h.userType.Sex != "其他" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sex"})
		return
	}
	h.userType.CreatedAt = time.Now()
	h.userType.UpdatedAt = time.Now()

	result := h.db.Create(h.userType)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error, "message": "数据创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": true, "Message": "创建成功"})
}
