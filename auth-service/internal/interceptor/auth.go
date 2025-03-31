package interceptor

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
	authv1 "github.com/tclutin/article-alchemy-service-protos/gen/go/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	jwtManager manager.Manager
	methods    map[string]bool
}

func NewAuthInterceptor(jwtManager manager.Manager) *AuthInterceptor {
	methods := map[string]bool{
		authv1.AuthService_SignUp_FullMethodName:       false,
		authv1.AuthService_LogIn_FullMethodName:        false,
		authv1.AuthService_RefreshToken_FullMethodName: false,
		authv1.AuthService_Logout_FullMethodName:       true,
		authv1.AuthService_GetUserInfo_FullMethodName:  true,
	}

	return &AuthInterceptor{
		jwtManager: jwtManager,
		methods:    methods,
	}
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if !a.methods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		tokenHeader := md["token"]
		if len(tokenHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization header is missing")
		}

		userId, err := a.jwtManager.ParseToken(tokenHeader[0])
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid access token")
		}

		ctx = context.WithValue(ctx, "userId", userId)

		return handler(ctx, req)
	}
}
