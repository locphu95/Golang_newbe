package transporthttp

import (
	"context"
	"log"
	"net/http"
	"time"
)

type AppHandler func(ctx context.Context, r *http.Request) (any, error)

func Execute(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		ctx := r.Context()

		// ===== call business handler =====
		resp, err := h(ctx, r)

		// ===== error handling =====
		if err != nil {
			writeError(w, err)
			return
		}

		// ===== success response =====
		writeJSON(w, resp)
		reqID, _ := r.Context().Value(RequestIDKey).(string)

		log.Printf("[%s] Request %s in %v", reqID, r.URL.Path, time.Since(start))
	}
}
