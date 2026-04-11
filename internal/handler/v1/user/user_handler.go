package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	service "github.com/locphu95/smart_machine/backend-core/internal/services"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
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

func (h *UserHandler) CreateUser(ctx context.Context,
	r *http.Request) (any, error) {
	id := chi.URLParam(r, "id")

	user, err := h.Service.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // lấy context từ request

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Register(ctx, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	})
}
