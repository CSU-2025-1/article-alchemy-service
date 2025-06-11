package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
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

	path := "./configs/config.yaml"

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		log.Fatalln("failed to read config:", err)
	}

	return &config
}
