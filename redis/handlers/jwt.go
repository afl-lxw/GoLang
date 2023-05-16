package handlers

import (
	"time"
)

// 定义 JWT 的有效期
var tokenExpiration = 24 * time.Hour

// 将 JWT 存储到 Redis
func storeTokenToRedis(userID string, token string) error {
	//err := redis.Redis.Set(userID, token, tokenExpiration).Err()
	//if err != nil {
	//	return err
	//}
	return nil
}
