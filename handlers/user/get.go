package user

import (
	UserType "Golang/models/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type UserHandler struct {
	db   *gorm.DB
	user *UserType.User
}

const layout = "2006-01-02 15:04:05"

func GetUserList(db *gorm.DB) *UserHandler {
	u := &UserType.User{}
	// 可以在这里进行依赖注入
	return &UserHandler{db: db, user: u}
}

// GetUserList @Summary 获取用户列表
// @Description 获取所有用户的列表数据
// @Tags 用户管理
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param keyword query string false "关键词"
// @Success 200 {object} 返回用户列表数据
// @Failure 500 {object} 返回错误信息
// @Router /users [get]
func (h *UserHandler) GetUserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	keyword := c.Query("keyword")

	// 查询总数
	var count int64
	query := h.db.Model(h.user)
	if keyword != "" {
		query = query.Where("Username LIKE ?", "%"+keyword+"%").Or("Mobile LIKE ?", "%"+keyword+"%")
	}
	query.Count(&count)

	// 分页查询
	var userList []*UserType.User
	//err := query.Offset((page - 1) * size).Limit(size).Find(&userList).Error
	// 分页查询用户列表，并只返回指定的字段
	err := query.Select("id, username, age, gender, mobile, create_at, update_at, isDelete").
		Limit(size).Offset((page - 1) * size).Find(&userList).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "查询用户列表失败",
			"data":    nil,
		})
		return
	}

	var pageNext bool
	if int64(page*size) < count {
		pageNext = true
	} else {
		pageNext = false
	}

	c.JSON(200, gin.H{
		"message": "",
		"data": gin.H{
			"total":    count,
			"page":     page,
			"size":     size,
			"data":     userList,
			"pageNext": pageNext,
		},
	})
}
