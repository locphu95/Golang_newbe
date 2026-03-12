package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locphu95/smart_machine/backend-core/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) GetUser(ctx context.Context,
	r *http.Request) (any, error) {
	id := chi.URLParam(r, "id")

	user, err := h.Service.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
