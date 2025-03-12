package config

import (
	"log"
	"os"
	"strconv"
)

// Config структура для хранения переменных окружения
type Config struct {
	DbUrl             string
	Port              string
	JwtSecret         string
	JwtSecretLifeTime int
}

// NewConfig загружает переменные из .env и возвращает структуру Config
func NewConfig() *Config {
	secretLifeTime, err := strconv.Atoi(os.Getenv("JWT_SECRET_LIFE_TIME"))
	if err != nil {
		log.Panic("Could not conver jwt secret lifetime: ", err)
	}

	return &Config{
		DbUrl:             os.Getenv("DATABASE_URL"),
		Port:              os.Getenv("PORT"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		JwtSecretLifeTime: secretLifeTime,
	}
}
