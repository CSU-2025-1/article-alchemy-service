package app

import (
	"context"
	"github.com/CSU-2025-1/article-alchemy-service/api_gateway/internal/config"
	"github.com/CSU-2025-1/article-alchemy-service/api_gateway/internal/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tclutin/article-alchemy-service-protos/gen/go/auth_v1"
	"github.com/tclutin/article-alchemy-service-protos/gen/go/content_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	cfg := config.MustLoadConfig()

	authAddrs := cfg.Services["auth_service"].Instances
	contentAddrs := cfg.Services["content_service"].Instances

	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	mux := runtime.NewServeMux(
		runtime.WithMetadata(middleware.AttachTokenMetadata),
		runtime.WithMetadata(middleware.AttachIpAddressMetadata),
	)

	err := auth_v1.RegisterAuthServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		authAddrs[0],
		grpcOpts)

	if err != nil {
		panic(err)
	}

	err = content_v1.RegisterContentServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		contentAddrs[0],
		grpcOpts)

	if err != nil {
		panic(err)
	}

	corsMux := middleware.AllowCORSMiddleware(mux)

	err = http.ListenAndServe(net.JoinHostPort(cfg.HTTPServer.Host, cfg.HTTPServer.Port), corsMux)
	if err != nil {
		panic(err)
	}

}
