package interceptor

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
)

type AuthInterceptor struct {
	jwtManager manager.Manager
	methods    map[string]bool
}

func NewAuthInterceptor(jwtManager manager.Manager, methods map[string]bool) *AuthInterceptor {
	return &AuthInterceptor{jwtManager: jwtManager, methods: methods}
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if !a.methods[info.FullMethod] {
			slog.Info("method", info.FullMethod)
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
