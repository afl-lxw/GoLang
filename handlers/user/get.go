package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	db   *sql.DB
	user *Usertype
}

const layout = "2006-01-02 15:04:05"

func GetUserList(db *sql.DB) *UserHandler {
	u := &Usertype{}
	// 可以在这里进行依赖注入
	return &UserHandler{db: db, user: u}
}

func (h *UserHandler) GetUserList(c *gin.Context) {
	// 执行 SQL 语句，获取用户列表数据
	rows, err := h.db.Query("SELECT Id, Username, Age,Gender,Mobile,create_at,update_at,User_Id,Password,IsDelete,Salt FROM user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取用户列表失败",
			"error":   err.Error(),
		})
		return
	}
	defer rows.Close()

	// 将查询结果保存到 users 数组中
	var users []map[string]interface{}
	var createdAtStr, updatedAtStr, UserId string
	for rows.Next() {
		err := rows.Scan(
			&h.user.Id,
			&h.user.Username,
			&h.user.Age,
			&h.user.Gender,
			&h.user.Mobile,
			&createdAtStr,
			&updatedAtStr,
			&UserId,
			&h.user.Password,
			&h.user.IsDelete,
			&h.user.Salt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "获取用户列表失败~",
				"error":   err.Error(),
			})
			return
		}
		createdAt, err := time.Parse(layout, createdAtStr)
		updatedAt, err := time.Parse(layout, updatedAtStr)
		h.user.CreateAt = createdAt
		h.user.UpdateAt = updatedAt
		user := map[string]interface{}{
			"id":         h.user.Id,
			"username":   h.user.Username,
			"age":        h.user.Age,
			"gender":     h.user.Gender,
			"mobile":     h.user.Mobile,
			"created_at": h.user.CreateAt,
			"updated_at": h.user.UpdateAt,
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		// 处理查询出错的情况
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取用户列表失败---",
			"error":   err.Error(),
		})
		return
	}

	// 返回用户列表数据
	c.JSON(http.StatusOK, gin.H{
		"message": "获取用户列表成功",
		"data":    users,
	})
}
