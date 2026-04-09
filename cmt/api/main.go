package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	transporthttp "github.com/locphu95/smart_machine/backend-core/internal/transport/http/middleware"
	routers "github.com/locphu95/smart_machine/backend-core/internal/transport/http/routers"
	"github.com/locphu95/smart_machine/backend-core/pkg/config"
	db "github.com/locphu95/smart_machine/backend-core/pkg/database"
	"github.com/locphu95/smart_machine/backend-core/pkg/redis"
	"github.com/locphu95/smart_machine/backend-core/pkg/server"
)

func main() {
	// 1. Load config
	cfg := config.Load()

	// Init Postgres
	pg := db.NewPostgres(cfg.PostgresConn)
	if err := pg.Connect(); err != nil {
		log.Fatalf("Postgres connect error: %v", err)
	}
	defer pg.Close()

	// Init Redis
	redis := redis.NewRedis(cfg.RedisHost, cfg.RedisPort)
	if err := redis.Connect(); err != nil {
		log.Fatalf("Redis connect error: %v", err)
	}
	defer redis.Close()

	// 2. Init router
	log.Printf("-->>> load router <<<--")
	r := chi.NewRouter()
	r.Use(transporthttp.RequestID)
	r.Use(transporthttp.Logger)
	r.Use(transporthttp.Recovery)
	// 3. Register routes
	routers.RegisterRoutes(r)

	// 4. Start server
	log.Printf("-->>> start server <<<--")

	addr := ":8080"
	if cfg.AppPort != "" {
		addr = ":" + cfg.AppPort
	}
	s := server.New(addr, r)
	log.Printf("Server running at %s\n", addr)
	if err := s.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
