package config

import (
	"time"
)

type (
	Redis struct {
		Host        string
		Password    string
		Database    int
		IdleTimeout time.Duration
	}
)

type (
	Configure struct {
		Redis *Redis
	}
)
