package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func InitRedis() (*Redis, error) {
	addr := viper.GetString("redis.host")
	redisPort := viper.GetInt("redis.port")
	addr = addr + ":" + strconv.Itoa(redisPort)
	redisPassword := viper.GetString("redis.password")
	redisDB := viper.GetInt("redis.db")
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     redisPassword,
		DB:           redisDB,
		PoolSize:     10,
		MinIdleConns: 5,
		IdleTimeout:  10 * time.Minute,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.New("failed to connect to Redis")
	}
	fmt.Println("Successfully connected to the Redis!")
	r := &Redis{Client: client}
	return r, nil
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
	fmt.Println("key--->:", key, "value:", value, "expiration:", expiration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("ctx:", ctx)
	fmt.Println("r.Client:", r.Client)
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
