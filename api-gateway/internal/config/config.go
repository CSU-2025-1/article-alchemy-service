package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	HTTPServer HTTPServer         `yaml:"http_server"`
	Services   map[string]Service `yaml:"services"`
}

type HTTPServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Service struct {
	Instances []string `yaml:"instances"`
}

func MustLoadConfig() *Config {
	var config Config

	path := os.Getenv("CONFIG_PATH_GATEWAY")
	if len(path) == 0 {
		log.Fatalln("CONFIG_PATH_GATEWAY environment variable not set")
	}

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		log.Fatalln("failed to read config:", err)
	}

	return &config
}
