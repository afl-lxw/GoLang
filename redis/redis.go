package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// 全局定义
var (
	RedisDb *redis.Client
)

// 创建 redis 链接
func init() {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "111.229.91.20:6379",
		Password: "chengqiang", // no password set
		DB:       0,            // use default DB
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		//连接失败
		println(err)
	}
}
