package grpc

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/converter"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain"

	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/auth"
	authv1 "github.com/tclutin/article-alchemy-service-protos/gen/go/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(grpcServer *grpc.Server) {
	authv1.RegisterAuthServiceServer(grpcServer, h)
}

func (h *AuthHandler) SignUp(ctx context.Context, request *authv1.RegisterRequest) (*authv1.TokenResponse, error) {
	tokens, err := h.authService.SignUp(ctx, auth.SignUpDTO{
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		return nil, converter.ConvertErrorToGRPCStatus(err)
	}

	return &authv1.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *AuthHandler) LogIn(ctx context.Context, request *authv1.LoginRequest) (*authv1.TokenResponse, error) {
	tokens, err := h.authService.LogIn(ctx, auth.LogInDTO{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return nil, converter.ConvertErrorToGRPCStatus(err)
	}

	return &authv1.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *AuthHandler) RefreshToken(ctx context.Context, request *authv1.RefreshTokenRequest) (*authv1.TokenResponse, error) {
	tokens, err := h.authService.RefreshToken(ctx, request.RefreshToken)

	if err != nil {
		return nil, converter.ConvertErrorToGRPCStatus(err)
	}

	return &authv1.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil

}

func (h *AuthHandler) GetUserInfo(ctx context.Context, _ *emptypb.Empty) (*authv1.GetUserInfoResponse, error) {
	userId := ctx.Value("userId")
	if userId == nil {
		return nil, status.Error(codes.PermissionDenied, "User not found in context")
	}

	usr, err := h.authService.GetUserInfo(ctx, userId.(uint64))

	if err != nil {
		return nil, converter.ConvertErrorToGRPCStatus(err)
	}

	return &authv1.GetUserInfoResponse{
		UserId:    usr.UserID,
		Username:  usr.Username,
		Email:     usr.Email,
		Password:  usr.PasswordHash,
		CreatedAt: timestamppb.New(usr.CreatedAt),
	}, nil
}

func (h *AuthHandler) Logout(ctx context.Context, request *authv1.LogoutRequest) (*authv1.LogoutResponse, error) {
	if err := h.authService.Logout(ctx, request.RefreshToken); err != nil {
		return nil, converter.ConvertErrorToGRPCStatus(err)
	}

	return &authv1.LogoutResponse{
		Message: "Logout Success",
	}, nil
}
