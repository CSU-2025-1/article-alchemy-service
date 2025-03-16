package handler

import (
	"context"
	"errors"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/auth"
	authpb "github.com/tclutin/article-alchemy-service-protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServer
	authService auth.Service
}

func NewAuthHandler(authService *auth.Service) *AuthHandler {
	return &AuthHandler{
		authService: *authService,
	}
}

func (h *AuthHandler) Register(grpcServer *grpc.Server) {
	authpb.RegisterAuthServer(grpcServer, h)
}

func (h *AuthHandler) SignUp(ctx context.Context, request *authpb.RegisterRequest) (*authpb.TokenResponse, error) {
	tokens, err := h.authService.SignUp(ctx, auth.SignUpDTO{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *AuthHandler) LogIn(ctx context.Context, request *authpb.LoginRequest) (*authpb.TokenResponse, error) {
	tokens, err := h.authService.LogIn(ctx, auth.LogInDTO{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		if errors.Is(err, domain.ErrWrongPassword) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *AuthHandler) RefreshToken(ctx context.Context, request *authpb.RefreshTokenRequest) (*authpb.TokenResponse, error) {
	tokens, err := h.authService.RefreshToken(ctx, request.RefreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrRefreshTokenNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil

}

func (h *AuthHandler) GetUserInfo(ctx context.Context, request *authpb.GetUserInfoRequest) (*authpb.UserInfoResponse, error) {
	usr, err := h.authService.GetUserByInfo(ctx, request.UserId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.UserInfoResponse{
		UserId:    usr.UserID,
		Username:  usr.Username,
		Email:     usr.Email,
		Password:  usr.PasswordHash,
		CreatedAt: timestamppb.New(usr.CreatedAt),
	}, nil
}

func (h *AuthHandler) Logout(ctx context.Context, request *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	if err := h.authService.Logout(ctx, request.UserId, request.RefreshToken); err != nil {
		if errors.Is(err, domain.ErrRefreshTokenNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.LogoutResponse{
		Message: "Logout Success",
	}, nil
}
