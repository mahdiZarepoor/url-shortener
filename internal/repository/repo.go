package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const expiration = time.Hour * 12

func InitializeStore() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Println("Redis started successfully")
	return redisClient
} 

func RegisterUrl(redisClient *redis.Client, shortUrl, originalUrl string) error{
	err := redisClient.Set(context.Background(), shortUrl, originalUrl, expiration).Err()
	if err != nil {
		log.Println(err.Error())
		return errors.New(fmt.Sprintf("error Set redis : %s", err.Error()))
	}
	return nil
}

func RetrieveUrl(redisClient *redis.Client, shortUrl string) (string, error) {
	url, err := redisClient.Get(context.Background(), shortUrl).Result()
	if err != nil {
		log.Println(err.Error())
		return "", errors.New(fmt.Sprintf("error Get redis : %s", err.Error()))
	}
	return url, nil
}
