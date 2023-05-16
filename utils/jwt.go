package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// 定义 JWT 签名密钥
var signingKey = []byte("Golang-signing-key")

// 定义 JWT 的有效期
var tokenExpiration = 24 * time.Hour

// GenerateJWT 生成 JWT
func GenerateJWT(userID uuid.UUID) (string, error) {
	// 创建 JWT 的声明
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(tokenExpiration).Unix(),
	}

	// 创建 Token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名 Token
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT 验证 JWT
func ValidateJWT(tokenString string) (string, error) {
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名密钥
		return signingKey, nil
	})
	if err != nil {
		return "", err
	}

	// 验证 Token 的有效性
	if token.Valid {
		// 获取用户 ID
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		return userID, nil
	}

	return "", fmt.Errorf("Invalid token")
}

// ProtectedHandler 刷新 JWT
func ProtectedHandler(c *gin.Context) {
	// 获取 Authorization 头部信息
	tokenString := c.GetHeader("Authorization")
	// 解析 JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 检查 JWT Token 是否有效
	if _, ok := token.Claims.(jwt.Claims); !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 从 Token 中提取用户信息等其他需要的数据
	//claims, ok := token.Claims.(jwt.MapClaims)
	//if !ok {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//	return
	//}

	// 在缓存中查找 JWT 过期时间
	//expirationTime, err := redisDB.Get(c, tokenString).Result()
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	//	return
	//}

	// 刷新 JWT 过期时间
	//newExpirationTime := time.Now().Add(30 * time.Minute)      // 更新为新的过期时间
	//redisDB.Set(c, tokenString, newExpirationTime.String(), 0) // 更新缓存中的过期时间

	// 返回响应给客户端
	return
}
