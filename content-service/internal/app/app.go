package app

import (
	"context"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/config"
	grpcHandler "github.com/CSU-2025-1/article-alchemy-service/content_service/internal/delivery/grpc"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/delivery/rabbitmq"
	contentService "github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/service"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/interceptor"
	contentRepository "github.com/CSU-2025-1/article-alchemy-service/content_service/internal/repository/content"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/pkg/client/postgresql"
	rabbitmqClient "github.com/CSU-2025-1/article-alchemy-service/content_service/pkg/client/rabbitmq"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/pkg/client/redis"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/pkg/jwt/manager"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/pkg/migrator"
	"github.com/rabbitmq/amqp091-go"
	"time"

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
	rabbitmq   *amqp091.Connection
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
	fmt.Println(jwtManager.NewAccessToken(1999, 5*time.Minute))

	conn, channel := rabbitmqClient.NewRabbitMQ(
		cfg.RabbitMQ.URL,
		cfg.RabbitMQ.ContentExchange,
		cfg.RabbitMQ.RequestContentParsing,
		cfg.RabbitMQ.ResponseContentParsing,
		cfg.RabbitMQ.Notification,
	)

	pub := rabbitmq.NewPublisher(channel, cfg.RabbitMQ.ContentExchange)

	cons := rabbitmq.NewConsumer(channel)

	contentRepo := contentRepository.NewContentRepository(postgresClient)

	contentSrv := contentService.NewService(contentRepo, pub, cons)

	go contentSrv.StartSomeShitListener(context.Background())

	contentHandler := grpcHandler.NewContentHandler(contentSrv)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.NewAuthInterceptor(jwtManager).Unary(),
			interceptor.ValidateInterceptor,
			interceptor.NewLimiterInterceptor(redisClient).Unary(),
		),
		grpc.Creds(insecure.NewCredentials()),
	)

	contentHandler.Register(grpcServer)

	return &App{
		pool:       postgresClient,
		rabbitmq:   conn,
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
	a.rabbitmq.Close()
}
