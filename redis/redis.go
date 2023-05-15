package redis

import (
	"Golang/config"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
}

func InitRedis(cfg *config.RedisConfig) (*Redis, error) {

	Client := redis.NewClient(&redis.Options{
		Addr:         cfg.Host,
		Password:     cfg.Password,
		DB:           cfg.Database,
		PoolSize:     10,
		MinIdleConns: 5,
		IdleTimeout:  cfg.IdleTimeout,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		return nil, errors.New("failed to connect to Redis")
	}
	fmt.Println("Successfully connected to the Redis!")

	return &Redis{Client: Client}, nil
}

func (r *Redis) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New("key does not exist")
		}
		return "", err
	}

	return val, nil
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Client.Del(ctx, key).Err()
}

func (r *Redis) Exists(key string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := r.Client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return val == 1, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
