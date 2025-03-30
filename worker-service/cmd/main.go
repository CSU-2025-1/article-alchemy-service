package main

import (
	"log"

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
	worker.StartWorker()
}
