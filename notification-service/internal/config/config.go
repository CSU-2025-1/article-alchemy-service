package config

import (
    "log"
    "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    RabbitMQ RabbitMQ
    SMTP     SMTP
}

type RabbitMQ struct {
    URL       string `env:"RABBITMQ_URL"`
    QueueName string `env:"RABBITMQ_QUEUE"`
}

type SMTP struct {
    Host     string `env:"SMTP_HOST"`
    Port     int    `env:"SMTP_PORT"`
    Username string `env:"SMTP_USER"`
    Password string `env:"SMTP_PASS"`
}

func MustLoad() *Config {
	var config Config

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatalln("error loading .env file")
	}

	return &config
}