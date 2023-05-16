package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTMiddleware 定义 JWT 中间件函数
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Token
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// 解析 Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名密钥
			return []byte("your-signing-key"), nil
		})

		if err != nil {
			// Token 解析失败，返回错误响应
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if token.Valid {
			// Token 验证通过，将解析后的用户信息保存到上下文中
			claims := token.Claims.(jwt.MapClaims)
			c.Set("user", claims["user"])
			c.Next()
		} else {
			// Token 验证失败，返回错误响应
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
