package config

import (
	"log"
	"os"
)

type Config struct {
	AppPort   string
	RedisHost string
	RedisPort string
	JwtSecret string
	AppName   string
}

func Load() *Config {
	log.Printf("-->>> load config <<<--")

	return &Config{
		AppPort:   os.Getenv("APP_PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		AppName:   os.Getenv("APP_NAME"),
	}
}
