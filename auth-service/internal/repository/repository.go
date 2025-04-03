package repository

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/user"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, usr user.User) (uint64, error)
	GetById(ctx context.Context, userId uint64) (user.User, error)
	GetByEmail(ctx context.Context, email string) (user.User, error)
}

type TokenRepository interface {
	Set(ctx context.Context, userID uint64, refreshToken string, ttl time.Duration) error
	Get(ctx context.Context, refreshToken string) (uint64, error)
	Del(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, userID uint64, oldRefreshToke string, refreshToken string, ttl time.Duration) error
}
