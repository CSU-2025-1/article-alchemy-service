package auth

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/repository"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
)

type Service struct {
	JWTManager manager.Manager
	TokenRepo  *repository.TokenRepository
	UserRepo   *repository.UserRepository
}

func NewService(
	jwtManager manager.Manager,
	userRepo *repository.UserRepository,
	tokenRepo *repository.TokenRepository,
) *Service {
	return &Service{
		JWTManager: jwtManager,
		TokenRepo:  tokenRepo,
		UserRepo:   userRepo,
	}
}

func (s *Service) SignUp(ctx context.Context, dto SignUpDTO) (TokenDTO, error) {
	// проверка если чел в бд с таким же email
	// если человек есть, то вкидываем ошиб-очку
	// если человека, то dto->usermodel
	// создание записи в бд
	// генерация refresh/access
	// сохранение рефреша по userid и deviceid
	//заебок
}

func (s *Service) LogIn(ctx context.Context, dto LogInDTO) (uint64, error) {
	panic("implement me")
}

func (s *Service) ValidateToken(ctx context.Context, token string) (uint64, error) {
	panic("implement me")
}

func (s *Service) Refresh(ctx context.Context, token string) (uint64, error) {
	panic("implement me")
}

func (s *Service) GetUserInfo(ctx context.Context, userID uint64) (uint64, error) {
	panic("implement me")
}
