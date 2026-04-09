package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locphu95/smart_machine/backend-core/internal/domain"
	service "github.com/locphu95/smart_machine/backend-core/internal/services"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUserV2(ctx context.Context,
	r *http.Request) (any, error) {
	id := chi.URLParam(r, "id")

	user, err := h.Service.GetUserV2(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}
