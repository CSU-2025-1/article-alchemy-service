package app

import (
	"context"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/config"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/pkg/client/postgresql"
)

type App struct {
	cfg  *config.Config
	pool postgresql.Client
}

func New() *App {

	// Init config
	cfg := config.MustLoad()

	dsn := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database)

	pool := postgresql.NewClient(context.Background(), dsn)

	return &App{
		cfg:  cfg,
		pool: pool,
	}
}

func (a *App) Run() {
	fmt.Println(a.cfg)
}
