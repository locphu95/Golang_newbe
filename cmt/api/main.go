package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	transporthttp "github.com/locphu95/smart_machine/backend-core/internal/transport/http/middleware"
	routers "github.com/locphu95/smart_machine/backend-core/internal/transport/http/routers"
	"github.com/locphu95/smart_machine/backend-core/pkg/config"
)

func main() {
	// 1. Load config
	cfg := config.Load()

	// 2. Init router
	log.Printf("-->>> load router <<<--")
	r := chi.NewRouter()
	r.Use(transporthttp.RequestID)
	r.Use(transporthttp.Logger)
	r.Use(transporthttp.Recovery)
	// 3. Register routes
	routers.RegisterRoutes(r)

	/* Newbee
	svc := &service.UserService{}
	handler := &handler.UserHandler{Service: svc}
	http.HandleFunc("/user", handler.GetUser)

	repo := &repository.UserReponsitoryImpl{}
	svc2 := service.NewUserService(repo)
	handler2 := &handler.UserHandler{Service: svc2}
	http.HandleFunc("/user2", handler2.GetUser)
	*/
	// add logger

	// 4. Start server
	log.Printf("-->>> start server <<<--")

	addr := ":8080"
	if cfg.AppPort != "" {
		addr = ":" + cfg.AppPort
	}

	log.Printf("Server running at %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
