package server

import (
	"context"
	"encoding/json"
	"fmt"
	"itemmeli/models"
	"itemmeli/package/config"
	"log"
	"net/http"
	"strconv"
	"time"

	"itemmeli/metrics"

	"github.com/gorilla/mux"
)

type contextKey string

const pathKey contextKey = "path"

type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

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

func routePathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if route := mux.CurrentRoute(r); route != nil {
			if tmpl, err := route.GetPathTemplate(); err == nil {
				ctx := context.WithValue(r.Context(), pathKey, tmpl)
				r = r.WithContext(ctx)
			}
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
		sw := &statusResponseWriter{ResponseWriter: w, status: http.StatusOK}
		go func() {
			next.ServeHTTP(sw, r)
			close(done)
		}()

		select {
		case <-done:
			elapsed := time.Since(start)
			path := r.Context().Value("path").(string)
			metrics.UpdTimeResponse(r.Method, path, strconv.Itoa(sw.status), elapsed.Seconds())
			metrics.IncHttpRequestsTotal(r.Method, path, strconv.Itoa(sw.status))
			log.Printf("%s %s finished in %s", r.Method, r.URL.Path, elapsed)
			return

		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				resp := models.Response{
					Success: false,
					Status:  http.StatusRequestTimeout,
					Message: RequestTimedOut,
					Error:   RequestTimedOut,
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

	router.Use(routePathMiddleware, enableCORS)

	commonAddress := fmt.Sprintf("%s:%d", config.Host(), config.Port())
	return &http.Server{
		Addr:    commonAddress,
		Handler: router,
	}, router
}
