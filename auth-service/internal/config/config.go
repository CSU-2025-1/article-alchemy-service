package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	JWT        JWT
	Redis      Redis
	Postgres   Postgres
	GRPCServer GRPCServer
}

type JWT struct {
	Secret        string        `env:"JWT_SECRET"`
	AccessExpire  time.Duration `env:"JWT_ACCESS_EXPIRE"`
	RefreshExpire time.Duration `env:"JWT_REFRESH_EXPIRE"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DATABASE"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
}

type GRPCServer struct {
	Host string `env:"GRPC_HOST"`
	Port string `env:"GRPC_PORT"`
}

func MustLoad() *Config {
	var config Config

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Println("error loading .env file")
	}

	return &config
}
