package app

import (
	"context"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/config"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/auth"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/handler"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/interceptor"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/repository"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/client/postgresql"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/client/redis"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/jwt/manager"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/migrator"
	"github.com/jackc/pgx/v5/pgxpool"
	authpb "github.com/tclutin/article-alchemy-service-protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
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

	userRepo := repository.NewUserRepository(postgresClient)

	tokenRepo := repository.NewTokenRepository(redisClient)

	authService := auth.NewService(cfg.JWT, jwtManager, userRepo, tokenRepo)

	authHandler := handler.NewAuthHandler(authService)

	authInterceptor := interceptor.NewAuthInterceptor(jwtManager, map[string]bool{
		authpb.AuthService_SignUp_FullMethodName:       false,
		authpb.AuthService_LogIn_FullMethodName:        false,
		authpb.AuthService_RefreshToken_FullMethodName: false,
		authpb.AuthService_Logout_FullMethodName:       true,
		authpb.AuthService_GetUserInfo_FullMethodName:  true,
	})

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()), grpc.Creds(insecure.NewCredentials()))

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
		log.Println("pprof running on :6060")
		if err := http.ListenAndServe("localhost:1010", nil); err != nil {
			log.Printf("pprof failed: %v", err)
		}
	}()

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
