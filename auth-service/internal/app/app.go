package app

import (
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/config"
)

type App struct {
	cfg *config.Config
}

func New() *App {

	// Init config
	cfg := config.MustLoad()

	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() {
	fmt.Println(a.cfg)
}
