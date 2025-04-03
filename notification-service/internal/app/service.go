package app

import (
    "log"
    "github.com/CSU-2025-1/article-alchemy-service/notification-service/internal/config"
)

type Server struct {
    consumer *Consumer
    cfg      *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
    
    consumer, err := NewConsumer(cfg.RabbitMQ.URL, cfg.RabbitMQ.QueueName)
    if err != nil {
        return nil, err
    }
    
    return &Server{
        consumer: consumer,
        cfg:      cfg,
    }, nil
}

func (s *Server) Run() error {
    log.Println("Starting email notification service...")
    return s.consumer.StartConsuming()
}

func (s *Server) Shutdown() {
    s.consumer.Close()
}