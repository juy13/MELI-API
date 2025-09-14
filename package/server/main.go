package server

import (
	"context"
	"fmt"
	"itemmeli/package/config"
	"itemmeli/package/service"
	"net/http"
	"strconv"

	_ "itemmeli/docs/swagger"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type ServerV1 struct {
	meliService service.Service
	server      *http.Server
	router      *mux.Router

	info string
}

// @title MELI Item Detail
// @version 1.0
// @description This is a server to get item details for MercadoLibre frontend page.
// @termsOfService http://swagger.io/terms/

func NewServerV1(service service.Service, config config.APIConfig) *ServerV1 {
	muxServer, router := NewMuxServer(config)
	server := &ServerV1{
		meliService: service,
		server:      muxServer,
		info:        fmt.Sprintf("Running server on %v", config.Host()+":"+strconv.Itoa(config.Port())),
		router:      router,
	}
	server.registerRoutes()
	return server
}

func (s *ServerV1) Info() string {
	return s.info
}

func (s *ServerV1) registerRoutes() {
	s.router.HandleFunc("/api/v1/item", s.itemDetails).Methods("GET")
	s.router.HandleFunc("/api/v1/recommendations", s.recommendations).Methods("GET")
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // swagger.json route
	))
}

func (s *ServerV1) Start() error {
	return s.server.ListenAndServe()
}

func (s *ServerV1) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// @Summary		Get item info
// @Description	Getting information for provided item
// @Tags 		item
// @Param		itemID path integer 	true 	"The users id"
// @Param		userID query integer 	true 	"User id"
// @Success		200
// @Failure		400					"The specified URL is invalid"
// @Failure		404					"Not found"
// @Failure		408 				"Request timed out"
// @Failure		500					"Internal server error"
// @Router		/api/v1/item/{itemID} [get]
func (s *ServerV1) itemDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// @Summary		Get recommendations
// @Description	Get recommendations for a given user and item
// @Tags 		item
// @Param		itemID path integer 	true 	"The users id"
// @Param		userID query integer 	true 	"User id"
// @Success		200
// @Failure		400					"The specified URL is invalid"
// @Failure		404					"Not found"
// @Failure		408 				"Request timed out"
// @Failure		500					"Internal server error"
// @Router		/api/v1/recommendations/{itemID} [get]
func (s *ServerV1) recommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
