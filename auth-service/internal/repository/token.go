package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type TokenRepository struct {
	redisClient *redis.Client
	prefix      string
}

func NewTokenRepository(redis *redis.Client) *TokenRepository {
	return &TokenRepository{
		redisClient: redis,
		prefix:      "refresh_token",
	}
}

func (t *TokenRepository) Set(ctx context.Context, userID uint64, refreshToken string, ttl time.Duration) error {
	key := fmt.Sprintf("%s:%s", t.prefix, refreshToken)
	return t.redisClient.Set(ctx, key, userID, ttl).Err()
}

func (t *TokenRepository) Get(ctx context.Context, refreshToken string) (uint64, error) {
	key := fmt.Sprintf("%s:%s", t.prefix, refreshToken)

	userID, err := t.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	parseUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid userID in redis: %w", err)
	}

	return parseUint, nil
}

func (t *TokenRepository) Del(ctx context.Context, refreshToken string) error {
	key := fmt.Sprintf("%s:%s", t.prefix, refreshToken)
	return t.redisClient.Del(ctx, key).Err()
}
