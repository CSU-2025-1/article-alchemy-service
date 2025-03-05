package auth

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/service/user"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	Create(ctx context.Context, user user.User) (uint64, error)
	GetById(ctx context.Context, userID uint64) (user.User, error)
}

type Service struct {
	JWTManager manager.TokenManager
	Redis      *redis.Client
	UserRepo   Repository
}

func NewService(jwtManager manager.TokenManager, userRepo Repository, client redis.Client) *Service {
	return &Service{JWTManager: jwtManager, UserRepo: userRepo}
}

func (s *Service) SignUp(ctx context.Context, dto SignUpDTO) (uint64, error) {
	panic("implement me")
}

func (s *Service) LogIn(ctx context.Context, dto LogInDTO) (uint64, error) {
	panic("implement me")
}

func (s *Service) ValidateToken(ctx context.Context, jwtToken string) (uint64, error) {
	panic("implement me")
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (uint64, error) {
	panic("implement me")
}

func (s *AuthService) GetUserInfo(ctx context.Context, userID uint64) (uint64, error) {
	panic("implement me")
}
