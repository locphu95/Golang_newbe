package transporthttp

import (
	"errors"
	"net/http"

	"github.com/locphu95/smart_machine/backend-core/internal/domain"
)

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

func mapError(err error) (int, string) {

	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		return http.StatusNotFound, err.Error()

	case errors.Is(err, domain.ErrInvalidInput):
		return http.StatusBadRequest, err.Error()

	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
