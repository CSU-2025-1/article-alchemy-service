package app

import (
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/api_gateway/internal/config"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	cfg := config.MustLoadConfig()
	fmt.Println(cfg)

}
