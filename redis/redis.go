package redis

import (
	"Golang/config"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisStore 是 Redis 客户端的封装
type RedisStore struct {
	client *redis.Client
}

var _ *config.Configure

// NewRedisStore 返回一个 RedisStore 实例
func NewRedisStore(configs *config.Redis) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 设置 Redis 密码
		DB:       0,  // 使用默认的数据库
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		panic(err)
	}
	// 存储 Redis 配置信息到全局变量
	_ = &config.Configure{
		Redis: configs,
	}

	//config.Configure{  }
	return &RedisStore{client: rdb}
}

// Set 将一个键值对写入 Redis
func (rs *RedisStore) Set(key string, value interface{}, expiration time.Duration) error {
	err := rs.client.Set(rs.client.Context(), key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s: %w", key, err)
	}

	return nil
}

// Get 从 Redis 中获取指定键的值
func (rs *RedisStore) Get(key string) (string, error) {
	value, err := rs.client.Get(rs.client.Context(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist", key)
		}
		return "", fmt.Errorf("failed to get key %s: %w", key, err)
	}

	return value, nil
}

// Del 从 Redis 中删除指定键
func (rs *RedisStore) Del(key string) error {
	err := rs.client.Del(rs.client.Context(), key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key %s: %w", key, err)
	}

	return nil
}
