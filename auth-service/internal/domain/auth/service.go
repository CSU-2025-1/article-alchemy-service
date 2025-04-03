package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/config"
	domainErr "github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/errors"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/user"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/repository"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/hasher"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
	"github.com/jackc/pgx/v5"
	"time"
)

type Service struct {
	jwtConfig  config.JWT
	jwtManager manager.Manager
	userRepo   repository.UserRepository
	tokenRepo  repository.TokenRepository
}

func NewService(
	jwtConfig config.JWT,
	jwtManager manager.Manager,
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
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
		return TokenDTO{}, domainErr.ErrUserAlreadyExists
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
		if errors.Is(err, domainErr.ErrUserNotFound) {
			return TokenDTO{}, err
		}
		return TokenDTO{}, fmt.Errorf("error getting user: %w", err)
	}

	if !hasher.CompareBcryptHash(usr.PasswordHash, dto.Password) {
		return TokenDTO{}, domainErr.ErrWrongPassword
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
		return TokenDTO{}, domainErr.ErrRefreshTokenNotFound
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

func (s *Service) GetUserInfo(ctx context.Context, userId uint64) (user.User, error) {
	usr, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return usr, domainErr.ErrUserNotFound
		}
		return usr, fmt.Errorf("error getting user: %w", err)
	}

	return usr, nil
}

// Logout TODO: get access and move its blacklist
func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	userID, err := s.tokenRepo.Get(ctx, refreshToken)
	if err != nil {
		return fmt.Errorf("error getting refresh token: %w", err)
	}

	if userID == 0 {
		return domainErr.ErrRefreshTokenNotFound
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
