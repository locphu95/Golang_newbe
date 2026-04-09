package router

import (
	"github.com/go-chi/chi/v5"
	v2handler "github.com/locphu95/smart_machine/backend-core/internal/handler/v1.1/user"
	v1handler "github.com/locphu95/smart_machine/backend-core/internal/handler/v1/user"
	"github.com/locphu95/smart_machine/backend-core/internal/repository"
	service "github.com/locphu95/smart_machine/backend-core/internal/services"
	transporthttp "github.com/locphu95/smart_machine/backend-core/internal/transport/http/middleware"
)

func RegisterUserRoutes(r chi.Router) {

	// ===== Dependency Injection =====x
	repo := &repository.UserRepositoryImpl{}
	svc := service.NewUserService(repo)

	handlerV1 := &v1handler.UserHandler{Service: svc}
	handlerV2 := &v2handler.UserHandler{Service: svc}

	// ===== API V1 =====
	r.Route("/v1", func(r chi.Router) {
		r.Post("/{id}", transporthttp.Execute(handlerV1.GetUser))
	})

	// ===== API V2 =====
	r.Route("/v2", func(r chi.Router) {
		r.Post("/{id}", transporthttp.Execute(handlerV2.GetUserV2))
		r.Post("/create", transporthttp.Execute(handlerV2.GetUserV2))

	})

}
