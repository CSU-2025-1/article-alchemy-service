package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type TokenRepository struct {
	redisClient *redis.Client
	prefix      string
}

func NewTokenRepository(redis *redis.Client) *TokenRepository {
	return &TokenRepository{
		redisClient: redis,
		prefix:      "refresh_token:",
	}
}

func (t *TokenRepository) Set(ctx context.Context, userID uint64, deviceID string, refreshToken string, ttl time.Duration) error {
	key := fmt.Sprintf("%s%d:%s", t.prefix, userID, deviceID)
	return t.redisClient.Set(ctx, key, refreshToken, ttl).Err()
}

func (t *TokenRepository) Get(ctx context.Context, userID uint64, deviceID string) (string, error) {
	key := fmt.Sprintf("%s%d:%s", t.prefix, userID, deviceID)
	return t.redisClient.Get(ctx, key).Result()
}

func (t *TokenRepository) Delete(ctx context.Context, userID uint64, deviceID string) error {
	key := fmt.Sprintf("%s%d:%s", t.prefix, userID, deviceID)
	return t.redisClient.Del(ctx, key).Err()
}
