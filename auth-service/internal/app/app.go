package app

import (
	"context"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/config"
	authService "github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/auth"
	grpcHandler "github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/handler/grpc"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/interceptor"
	tokenRepository "github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/repository/token"
	userRepository "github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/repository/user"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/client/postgresql"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/client/redis"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/migrator"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	pool       *pgxpool.Pool
	grpcServer *grpc.Server
	cfg        *config.Config
}

func New() *App {
	cfg := config.MustLoad()

	dsn := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database)

	postgresClient := postgresql.NewClient(context.Background(), dsn)

	redisClient := redis.NewClient(cfg.Redis.Host, cfg.Redis.Port)

	migrator.Migrate(postgresClient)

	jwtManager := manager.MustLoadTokenManager(cfg.JWT.Secret)

	userRepo := userRepository.NewUserRepository(postgresClient)

	tokenRepo := tokenRepository.NewTokenRepository(redisClient)

	authSrv := authService.NewService(cfg.JWT, jwtManager, userRepo, tokenRepo)

	authHandler := grpcHandler.NewAuthHandler(authSrv)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.NewAuthInterceptor(jwtManager).Unary(),
			interceptor.ValidateInterceptor,
		),
		grpc.Creds(insecure.NewCredentials()),
	)

	authHandler.Register(grpcServer)

	return &App{
		pool:       postgresClient,
		grpcServer: grpcServer,
		cfg:        cfg,
	}
}

func (a *App) Run() {
	listener, err := net.Listen("tcp", net.JoinHostPort(a.cfg.GRPCServer.Host, a.cfg.GRPCServer.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err = a.grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	a.Stop()
}

func (a *App) Stop() {
	a.grpcServer.GracefulStop()
	a.pool.Close()
}
