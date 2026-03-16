package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	v2handler "github.com/locphu95/smart_machine/backend-core/internal/handler/v1.1/user"
	v1handler "github.com/locphu95/smart_machine/backend-core/internal/handler/v1/user"
	"github.com/locphu95/smart_machine/backend-core/pkg/config"

	"github.com/locphu95/smart_machine/backend-core/internal/repository"
	"github.com/locphu95/smart_machine/backend-core/internal/service"
	transporthttp "github.com/locphu95/smart_machine/backend-core/internal/transport/middleware/http"
)

func main() {
	config.Load()
	r := chi.NewRouter()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OKE"))
	})

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
	r.Use(transporthttp.RequestID)
	r.Use(transporthttp.Logger)
	r.Use(transporthttp.Recovery)
	// ===== Dependency Injection =====
	repo := &repository.UserRepositoryImpl{}
	svc := service.NewUserService(repo)

	handlerV1 := &v1handler.UserHandler{Service: svc}
	handlerV2 := &v2handler.UserHandler{Service: svc}

	// ===== API V1 =====
	r.Route("/v1", func(r chi.Router) {
		r.Get("/user/{id}", transporthttp.Execute(handlerV1.GetUser))
	})

	// ===== API V2 =====
	r.Route("/v2", func(r chi.Router) {
		r.Get("/user/{id}", transporthttp.Execute(handlerV2.GetUserV2))
	})
	log.Println("server running at :8080")
	http.ListenAndServe(":8080", r)
}
