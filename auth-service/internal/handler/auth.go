package handler

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/auth"
	authpb "github.com/tclutin/article-alchemy-service-protos/gen/go/auth"
	"google.golang.org/grpc"
)

type AuthHandler struct {
	authService auth.Service
	authpb.UnimplementedAuthServer
}

func NewAuthHandler(authService *auth.Service) *AuthHandler {
	return &AuthHandler{
		authService: *authService,
	}
}

func (h *AuthHandler) Register(grpcServer *grpc.Server) {
	authpb.RegisterAuthServer(grpcServer, h)
}

func (h *AuthHandler) SignUp(
	ctx context.Context,
	request *authpb.RegisterRequest,
) (*authpb.TokenResponse, error) {

	_, err := h.authService.SignUp(ctx, auth.SignUpDTO{
		Email:    request.GetEmail(),
		Username: request.GetUsername(),
		Password: request.GetPassword(),
		DeviceID: request.GetDeviceId(),
	})

	if err != nil {
		return nil, err
	}

	return nil, err
}

func (h *AuthHandler) LogIn(
	ctx context.Context,
	request *authpb.LoginRequest,
) (*authpb.TokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AuthHandler) RefreshToken(
	ctx context.Context,
	request *authpb.RefreshTokenRequest,
) (*authpb.TokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AuthHandler) ValidateToken(
	ctx context.Context,
	request *authpb.AccessTokenRequest,
) (*authpb.ValidationTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h *AuthHandler) GetUserInfo(
	ctx context.Context,
	request *authpb.GetUserInfoRequest,
) (*authpb.UserInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}
