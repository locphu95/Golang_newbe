package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locphu95/smart_machine/backend-core/cmt/routes"
	transporthttp "github.com/locphu95/smart_machine/backend-core/internal/transport/middleware/http"
	"github.com/locphu95/smart_machine/backend-core/pkg/config"
)

func main() {
	config.Load()
	r := chi.NewRouter()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OKE"))
	})
	r.Use(transporthttp.RequestID)
	r.Use(transporthttp.Logger)
	r.Use(transporthttp.Recovery)
	routes.UserRouters(r)

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

	log.Println("server running at :8080")
	http.ListenAndServe(":8080", r)
}
