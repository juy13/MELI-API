package server

import (
	"context"
	"encoding/json"
	"fmt"
	"itemmeli/models"
	"itemmeli/package/config"
	"itemmeli/package/service"
	"net/http"
	"strconv"
	"time"

	_ "itemmeli/docs/swagger"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const (
	InternalServerError          = "Internal Server Error"
	NoUserFound                  = "No user found in request"
	ErrorFetchingItemDetails     = "Error fetching item details"
	ErrorFetchingRecommendations = "Error fetching recommendations"
	NoRecommendationsFound       = "No recommendations found"
	RequestTimedOut              = "Request timed out"

	SuccessItemDetails     = "Item details retrieved successfully"
	SuccessRecommendations = "Recommendations retrieved successfully"
)

type ServerV1 struct {
	meliService service.Service
	server      *http.Server
	router      *mux.Router

	info string

	timeout time.Duration
}

// @title MELI Item Detail
// @version 1.0
// @description This is an API to get item details for MercadoLibre frontend page.
// @termsOfService http://swagger.io/terms/

func NewServerV1(service service.Service, config config.APIConfig) *ServerV1 {
	muxServer, router := NewMuxServer(config)
	server := &ServerV1{
		meliService: service,
		server:      muxServer,
		info:        fmt.Sprintf("Running server on %v", config.Host()+":"+strconv.Itoa(config.Port())),
		router:      router,
		timeout:     config.RequestTimeout(),
	}
	server.registerRoutes()
	return server
}

func (s *ServerV1) Info() string {
	return s.info
}

func (s *ServerV1) registerRoutes() {
	s.router.Handle("/api/v1/item/{itemID}", timeoutMiddleware(s.timeout, http.HandlerFunc(s.itemDetails))).Methods("GET")
	s.router.Handle("/api/v1/recommendations/{itemID}", timeoutMiddleware(s.timeout, http.HandlerFunc(s.recommendations))).Methods("GET")
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // swagger.json route
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
// @Param		itemID path string 	true "The item id"
// @Param		userID query string true "User id"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response "The specified URL is invalid"
// @Failure		404 {object} models.Response "Not found"
// @Failure		408 {object} models.Response "Request timed out"
// @Failure		500 {object} models.Response "Internal server error"
// @Router		/api/v1/item/{itemID} [get]
func (s *ServerV1) itemDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	itemID := vars["itemID"]
	userID := r.URL.Query().Get("userID")

	resp := models.Response{
		Success: false,
		Status:  http.StatusInternalServerError,
		Message: InternalServerError,
		Error:   "",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")

	if userID == "" {
		select {
		case <-ctx.Done():
			return
		default:
			resp.Status = http.StatusBadRequest
			resp.Error = NoUserFound
			resp.Message = NoUserFound
			w.WriteHeader(resp.Status)
			_ = json.NewEncoder(w).Encode(resp)
			return
		}
	}

	if !s.checker(w, &resp, s.meliService.IsValidItem, ctx, itemID, "item") {
		return
	}

	if !s.checker(w, &resp, s.meliService.IsValidUser, ctx, userID, "user") {
		return
	}

	data, err := s.meliService.GetItemDetails(ctx, itemID, userID)
	if err != nil {
		select {
		case <-ctx.Done():
			return
		default:
			resp.Status = http.StatusInternalServerError
			resp.Error = err.Error()
			resp.Message = ErrorFetchingItemDetails
			w.WriteHeader(resp.Status)
			_ = json.NewEncoder(w).Encode(resp)
			return
		}
	}

	select {
	case <-ctx.Done():
		return
	default:
		resp.Success = true
		resp.Status = http.StatusOK
		resp.Message = SuccessItemDetails
		resp.Data = data

		w.WriteHeader(resp.Status)
		_ = json.NewEncoder(w).Encode(resp)
	}

}

// @Summary		Get recommendations
// @Description	Get recommendations for a given user and item
// @Tags 		item
// @Param		itemID path string 	true "The item id"
// @Param		sellerID query string true "Seller id"
// @Param		userID query string true "User id"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response "The specified URL is invalid"
// @Failure		404 {object} models.Response "Not found"
// @Failure		408 {object} models.Response "Request timed out"
// @Failure		500 {object} models.Response "Internal server error"
// @Router		/api/v1/recommendations/{itemID} [get]
func (s *ServerV1) recommendations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	itemID := vars["itemID"]
	sellerID := r.URL.Query().Get("sellerID")
	userID := r.URL.Query().Get("userID")

	resp := models.Response{
		Success: false,
		Status:  http.StatusInternalServerError,
		Message: InternalServerError,
		Error:   "",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")

	if !s.checker(w, &resp, s.meliService.IsValidItem, ctx, itemID, "item") {
		return
	}

	if !s.checker(w, &resp, s.meliService.IsValidUser, ctx, userID, "user") {
		return
	}

	recs, err := s.meliService.GetItemRecommendations(ctx, itemID, sellerID, userID)
	if err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Error = err.Error()
		resp.Message = ErrorFetchingRecommendations
		w.WriteHeader(resp.Status)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	if len(recs) == 0 {
		resp.Status = http.StatusNotFound
		resp.Message = NoRecommendationsFound
		w.WriteHeader(resp.Status)
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Success = true
	resp.Status = http.StatusOK
	resp.Message = SuccessRecommendations
	resp.Data = recs

	w.WriteHeader(resp.Status)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *ServerV1) checker(
	w http.ResponseWriter,
	resp *models.Response,
	validateFunc func(context.Context, string) (bool, error),
	ctx context.Context,
	id string,
	entityName string,
) bool {
	valid, err := validateFunc(ctx, id)
	if err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Error = err.Error()
		resp.Message = fmt.Sprintf("Error validating %s", entityName)
		w.WriteHeader(resp.Status)
		_ = json.NewEncoder(w).Encode(resp)
		return false
	}
	if !valid {
		resp.Status = http.StatusNotFound
		resp.Message = fmt.Sprintf("%s not found", entityName)
		w.WriteHeader(resp.Status)
		_ = json.NewEncoder(w).Encode(resp)
		return false
	}
	return true
}
