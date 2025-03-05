package handler

import (
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/auth"
	authpb "github.com/tclutin/article-alchemy-service-protos/gen/go/auth"
	"google.golang.org/grpc"
)

type Handler struct {
	//authpb.UnimplementedAuthServer
	authService auth.Service
}

func NewHandler(authService *auth.Service) *Handler {
	return &Handler{
		authService: *authService,
	}
}

func (h *Handler) Register(grpcServer *grpc.Server) {
	authpb.RegisterAuthServer(grpcServer, h)
}
