package config

import (
	"os"
)

// Config структура для хранения переменных окружения
type Config struct {
	DbUrl string `env:"DATABASE_URL"`
	Port  string `env:"PORT"`
}

// NewConfig загружает переменные из .env и возвращает структуру Config
func NewConfig() *Config {
	return &Config{
		DbUrl: os.Getenv("DATABASE_URL"),
		Port:  os.Getenv("PORT"),
	}
}
