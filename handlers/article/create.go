package article

import (
	"Golang/config"
	"Golang/models/article"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type ArticleType struct {
	db      *gorm.DB
	config  *config.Configure
	article *article.ArticleField
}

func NewArticle(db *gorm.DB) *ArticleType {
	redisClient := &config.Configure{
		Redis:       &config.RedisConfig{},
		RedisClient: &config.Redis{},
	}
	d := &article.ArticleField{}
	return &ArticleType{db: db, config: redisClient, article: d}
}

func (h *ArticleType) CreateArticle(c *gin.Context) {
	err := c.ShouldBindJSON(h.article)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	fmt.Printf("%+v", h.article.Title)
	fmt.Printf("%+v", h.article.Content)

	result := h.db.Create(h.article)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error, "message": "数据创建失败"})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
}
