package config

import (
	Selfredis "Golang/redis"
	"time"
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
		Client *Selfredis.Redis
	}
)

type (
	Configure struct {
		Redis       *RedisConfig
		RedisClient *Redis
	}
)
