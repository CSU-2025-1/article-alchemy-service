package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/config"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/user"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/hasher"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
	"github.com/jackc/pgx/v5"
	"time"
)

type UserRepo interface {
	Create(ctx context.Context, user user.User) (uint64, error)
	GetById(ctx context.Context, userID uint64) (user.User, error)
	GetByEmail(ctx context.Context, email string) (user.User, error)
}

type TokenRepo interface {
	Set(ctx context.Context, userID uint64, refreshToken string, ttl time.Duration) error
	Get(ctx context.Context, refreshToken string) (uint64, error)
	Del(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, userID uint64, oldRefreshToken string, refreshToken string, ttl time.Duration) error
}

type Service struct {
	jwtConfig  config.JWT
	jwtManager manager.Manager
	userRepo   UserRepo
	tokenRepo  TokenRepo
}

func NewService(
	jwtConfig config.JWT,
	jwtManager manager.Manager,
	userRepo UserRepo,
	tokenRepo TokenRepo,
) *Service {
	return &Service{
		jwtConfig:  jwtConfig,
		jwtManager: jwtManager,
		userRepo:   userRepo,
		tokenRepo:  tokenRepo,
	}
}

func (s *Service) SignUp(ctx context.Context, dto SignUpDTO) (TokenDTO, error) {
	_, err := s.userRepo.GetByEmail(ctx, dto.Email)
	if err == nil {
		return TokenDTO{}, domain.ErrUserAlreadyExists
	}

	hash, err := hasher.NewBcryptHash(dto.Password)
	if err != nil {
		return TokenDTO{}, fmt.Errorf("error hashing password: %w", err)
	}

	userID, err := s.userRepo.Create(ctx, user.User{
		Username:     dto.Username,
		Email:        dto.Email,
		IsActive:     false,
		PasswordHash: hash,
		CreatedAt:    time.Now().UTC(),
	})

	if err != nil {
		return TokenDTO{}, fmt.Errorf("error creating user: %w", err)
	}

	tokens, err := s.GenPairTokens(userID, s.jwtConfig.AccessExpire)
	if err != nil {
		return TokenDTO{}, fmt.Errorf("error generating pair tokens: %w", err)
	}

	if err = s.tokenRepo.Set(ctx, userID, tokens.RefreshToken, s.jwtConfig.RefreshExpire); err != nil {
		return TokenDTO{}, fmt.Errorf("error setting refresh token: %w", err)
	}

	return TokenDTO{
		tokens.AccessToken,
		tokens.RefreshToken,
	}, nil
}

func (s *Service) LogIn(ctx context.Context, dto LogInDTO) (TokenDTO, error) {
	usr, err := s.userRepo.GetByEmail(ctx, dto.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return TokenDTO{}, err
		}
		return TokenDTO{}, fmt.Errorf("error getting user: %w", err)
	}

	if !hasher.CompareBcryptHash(usr.PasswordHash, dto.Password) {
		return TokenDTO{}, domain.ErrWrongPassword
	}

	tokens, err := s.GenPairTokens(usr.UserID, s.jwtConfig.AccessExpire)
	if err != nil {
		return TokenDTO{}, fmt.Errorf("error generating pair tokens: %w", err)
	}

	if err = s.tokenRepo.Set(ctx, usr.UserID, tokens.RefreshToken, s.jwtConfig.RefreshExpire); err != nil {
		return TokenDTO{}, fmt.Errorf("error setting refresh token: %w", err)
	}

	return TokenDTO{
		tokens.AccessToken,
		tokens.RefreshToken,
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (TokenDTO, error) {
	userID, err := s.tokenRepo.Get(ctx, refreshToken)
	if err != nil {
		return TokenDTO{}, fmt.Errorf("error getting refresh token: %w", err)
	}

	if userID == 0 {
		return TokenDTO{}, domain.ErrRefreshTokenNotFound
	}

	tokens, err := s.GenPairTokens(userID, s.jwtConfig.AccessExpire)
	if err != nil {
		return TokenDTO{}, fmt.Errorf("error generating pair tokens: %w", err)
	}

	if err = s.tokenRepo.Refresh(ctx, userID, refreshToken, tokens.RefreshToken, s.jwtConfig.RefreshExpire); err != nil {
		return TokenDTO{}, fmt.Errorf("error refreshing token: %w", err)
	}

	return TokenDTO{
		tokens.AccessToken,
		tokens.RefreshToken,
	}, nil
}

func (s *Service) GetUserByInfo(ctx context.Context, userID uint64) (user.User, error) {
	usr, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return usr, domain.ErrUserNotFound
		}
		return usr, fmt.Errorf("error getting user: %w", err)
	}

	return usr, nil
}

func (s *Service) Logout(ctx context.Context, userID uint64, refreshToken string) error {
	// userID - рудимент
	userID, err := s.tokenRepo.Get(ctx, refreshToken)
	if err != nil {
		return fmt.Errorf("error getting refresh token: %w", err)
	}

	if userID == 0 {
		return domain.ErrRefreshTokenNotFound
	}

	if err = s.tokenRepo.Del(ctx, refreshToken); err != nil {
		return fmt.Errorf("error deleting refresh token: %w", err)
	}

	return nil
}

func (s *Service) GenPairTokens(userID uint64, ttl time.Duration) (TokenDTO, error) {
	accessToken, err := s.jwtManager.NewAccessToken(userID, ttl)
	if err != nil {
		return TokenDTO{}, err
	}

	refreshToken := s.jwtManager.NewRefreshToken().String()

	return TokenDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

/*
	func (s *Service) ValidateToken(ctx context.Context, token string) (uint64, error) {
		userID, err := s.jwtManager.ParseToken(token)
		if err != nil {
			return 0, domain.ErrMissingCredentials
		}

		if _, err = s.userRepo.GetById(ctx, userID); err != nil {
			return 0, domain.ErrUserNotFound
		}

		return userID, nil
	}
*/
