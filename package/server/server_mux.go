package server

import (
	"context"
	"encoding/json"
	"fmt"
	"itemmeli/models"
	"itemmeli/package/config"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func timeoutMiddleware(limit time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), limit)
		defer cancel()

		r = r.WithContext(ctx)

		start := time.Now()

		done := make(chan struct{})
		go func() {
			next.ServeHTTP(w, r)
			close(done)
		}()

		select {
		case <-done:
			elapsed := time.Since(start)
			log.Printf("%s %s finished in %s", r.Method, r.URL.Path, elapsed)
			return

		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				resp := models.Response{
					Success: false,
					Status:  http.StatusRequestTimeout,
					Message: "Request timed out",
					Error:   "Exceeded time limit",
					Data:    nil,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusRequestTimeout)
				_ = json.NewEncoder(w).Encode(resp)
			}
			return
		}
	})
}

func NewMuxServer(config config.APIConfig) (*http.Server, *mux.Router) {
	router := mux.NewRouter()

	handler := timeoutMiddleware(config.RequestTimeout(), enableCORS(router))

	commonAddress := fmt.Sprintf("%s:%d", config.Host(), config.Port())
	return &http.Server{
		Addr:    commonAddress,
		Handler: handler,
	}, router
}
