package contract

import "net/http"

type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}
