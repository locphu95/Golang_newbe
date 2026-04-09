package transporthttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func writeError(w http.ResponseWriter, err error) {

	status, msg := mapError(err)

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error:   msg,
	})
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    data,
	})
}
