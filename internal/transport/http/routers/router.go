package router

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {

	// init handlers

	// versioning
	r.Route("/user", func(r chi.Router) {
		RegisterUserRoutes(r)
	})
}
