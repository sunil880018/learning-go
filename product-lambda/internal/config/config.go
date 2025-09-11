package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	PostgresURL string
	RedisAddr   string
}

func Load() *Config {
	_ = godotenv.Load(".env")

	return &Config{
		Env:         os.Getenv("ENV"),
		PostgresURL: os.Getenv("POSTGRES_URL"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
	}
}
