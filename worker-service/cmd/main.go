package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"article-alchemy-service/internal/worker"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func main() {
	log.Println("Start worker...")
	go worker.StartWorker()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("Received shutdown signal, closing worker...")
}
