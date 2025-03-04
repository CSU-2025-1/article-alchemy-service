package app

import (
	"context"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/config"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/client/postgresql"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/client/redis"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	pool *pgxpool.Pool
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

	pool := postgresql.NewClient(context.Background(), dsn)

	redi1 := redis.NewClient(cfg.Redis.Host, cfg.Redis.Port)

	fmt.Println(redi1)

	return &App{
		pool: pool,
	}
}

func (a *App) Run() {
}
