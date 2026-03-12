package server

import (
	"net/http"
	"os"
)

func StartHTTPS(handler http.Handler) error {
	srv := &http.Server{
		Addr:    ":8443",
		Handler: handler,
	}

	return srv.ListenAndServeTLS(
		os.Getenv("TLS_CERT"),
		os.Getenv("TLS_KEY"),
	)
}
