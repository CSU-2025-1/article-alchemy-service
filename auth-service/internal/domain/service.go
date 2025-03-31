package domain

import (
	"context"
	"errors"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/auth"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/user"
)

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrUserNotFound         = errors.New("user not found")
	ErrWrongPassword        = errors.New("wrong password")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
)

type AuthService interface {
	SignUp(ctx context.Context, dto auth.SignUpDTO) (auth.TokenDTO, error)
	LogIn(ctx context.Context, dto auth.LogInDTO) (auth.TokenDTO, error)
	RefreshToken(ctx context.Context, refreshToken string) (auth.TokenDTO, error)
	GetUserInfo(ctx context.Context, userId uint64) (user.User, error)
	Logout(ctx context.Context, refreshToken string) error
}
