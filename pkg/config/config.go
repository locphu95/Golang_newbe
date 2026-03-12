package config

import "os"

type Config struct {
	AppPort   string
	RedisHost string
	RedisPort string
	JwtSecret string
}

func Load() *Config {
	return &Config{
		AppPort:   os.Getenv("APP_PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
