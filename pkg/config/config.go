package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	RedisHost    string
	RedisPort    string
	JwtSecret    string
	AppName      string
	PostgresConn string
}

func Load() *Config {
	log.Printf("-->>> load config <<<--")
	// Load file .env (hoặc .env.example nếu muốn)
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found, using system env")
	}
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")

	pgConn := "postgres://" + pgUser + ":" + pgPass + "@" + pgHost + ":" + pgPort + "/" + pgDB + "?sslmode=disable"
	log.Println("connect %s", pgConn)
	return &Config{
		AppPort:      os.Getenv("APP_PORT"),
		RedisHost:    os.Getenv("REDIS_HOST"),
		RedisPort:    os.Getenv("REDIS_PORT"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
		AppName:      os.Getenv("APP_NAME"),
		PostgresConn: pgConn,
	}
}
