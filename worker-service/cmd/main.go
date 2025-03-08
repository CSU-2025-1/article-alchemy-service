package main

import (
	"log"
	"net/http"

	"article-alchemy-service/internal/handler"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

func main() {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
