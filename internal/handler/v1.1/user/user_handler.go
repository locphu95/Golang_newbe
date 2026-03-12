package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/locphu95/smart_machine/backend-core/internal/domain"
	"github.com/locphu95/smart_machine/backend-core/internal/service"
)

type UserHandler struct {
	Service *service.UserService
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
