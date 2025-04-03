package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/app"
	"github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/config"
)

func main() {
	// Загрузка конфигурации
	cfg := config.MustLoad()
    
    application, err := app.NewServer(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize application: %v", err)
    }

	// Запуск сервиса в горутине
	go func() {
		if err := application.Run(); err != nil {
			log.Fatalf("Application failed: %v", err)
		}
	}()

	// Ожидание сигналов завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	log.Println("Shutting down server...")
	application.Shutdown()
	log.Println("Server exited properly")
}