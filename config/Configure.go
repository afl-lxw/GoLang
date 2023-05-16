package config

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RedisClient *redis.Client
)

type (
	RedisConfig struct {
		Host                 string
		Password             string
		Database             int
		IdleTimeout          time.Duration
		APIServerTaskRecover string
		APIServerTaskPause   string
		APIServerTaskStop    string
	}

	Redis struct {
		Client *redis.Client
	}
)

type (
	Configure struct {
		Redis       *RedisConfig
		RedisClient *Redis
	}
)
