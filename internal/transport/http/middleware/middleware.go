package transporthttp

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type ctxKey string

const RequestIDKey ctxKey = "request_id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)

		ctx := context.WithValue(r.Context(), RequestIDKey, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		reqID, _ := r.Context().Value(RequestIDKey).(string)

		if reqID == "" {
			reqID = "unknown"
		}

		next.ServeHTTP(w, r)

		log.Printf("[%s] %s %s took %v",
			reqID,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)

	})
}
