package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHealthEndpoint(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OKE"))
	})

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Body.String() != "OKE" {
		t.Errorf("expected OKE, got %s", w.Body.String())
	}
}

func TestUserV1Endpoint(t *testing.T) {
	r := chi.NewRouter()

	req := httptest.NewRequest("GET", "/v1/user/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	t.Logf("Response body: %s", w.Body.String())

	// tuỳ logic handler, bạn kiểm tra body JSON
	if !strings.Contains(w.Body.String(), "1") {
		t.Errorf("expected response to contain user id 123, got %s", w.Body.String())
	}
}
