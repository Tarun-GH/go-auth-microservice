package utils

import (
	"time"

	"github.com/Tarun-GH/go-rest-microservice/internal/config"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(client *redis.Client) {
	redisClient = client
}

func GenerateRefresh(userID int) (string, error) {
	refreshToken := uuid.NewString()

	err := redisClient.Set(
		config.Ctx,
		refreshToken,
		userID,
		7*24*time.Hour,
	).Err()
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func GetUserIDFromRefreshToken(token string) (int, bool) {
	val, err := redisClient.Get(config.Ctx, token).Int()
	if err != nil {
		return 0, false
	}
	return val, true
}

func DeleteRefreshToken(r_token string) error {
	return redisClient.Del(config.Ctx, r_token).Err()
}
