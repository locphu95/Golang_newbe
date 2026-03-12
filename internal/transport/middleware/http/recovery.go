package transporthttp

import "net/http"

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "internal error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
