package server

import (
	"context"
	"encoding/json"
	"itemmeli/mock"
	"itemmeli/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestItemDetails(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		mock       *mock.MockService
		wantStatus int
		wantMsg    string
	}{
		{
			name: "everything is ok",
			url:  "/api/v1/item/" + "item1" + "?userID=" + "user1",
			mock: &mock.MockService{
				IsValidItemF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				IsValidUserF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				GetItemF: func(ctx context.Context, itemID, client string) (*models.Item, error) {
					return &models.Item{ID: itemID, Title: "Test Item"}, nil
				},
			},
			wantStatus: http.StatusOK,
			wantMsg:    SuccessItemDetails,
		},
		{
			name: "invalid item",
			url:  "/api/v1/item/" + "bad-item" + "?userID=" + "user1",
			mock: &mock.MockService{
				IsValidItemF: func(ctx context.Context, id string) (bool, error) { return false, nil },
				IsValidUserF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				GetItemF:     nil,
			},
			wantStatus: http.StatusNotFound,
			wantMsg:    "item not found",
		},
		{
			name: "invalid user",
			url:  "/api/v1/item/" + "item1" + "?userID=" + "INVALID",
			mock: &mock.MockService{
				IsValidItemF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				IsValidUserF: func(ctx context.Context, id string) (bool, error) { return false, nil },
				GetItemF:     nil,
			},
			wantStatus: http.StatusNotFound,
			wantMsg:    "user not found",
		},
		{
			name: "timeout",
			url:  "/api/v1/item/" + "item1" + "?userID=" + "user1",
			mock: &mock.MockService{
				IsValidItemF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				IsValidUserF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				GetItemF: func(ctx context.Context, itemID, client string) (*models.Item, error) {
					// Simulate a long operation that exceeds deadline
					select {
					case <-ctx.Done():
						return nil, ctx.Err()
					case <-time.After(3 * time.Second):
						log.Print("Long operation completed successfully. Returning item.")
						return &models.Item{ID: itemID}, nil
					}
				},
			},
			wantStatus: http.StatusRequestTimeout,
			wantMsg:    RequestTimedOut,
		},
		{
			name: "no userid",
			url:  "/api/v1/item/" + "item1",
			mock: &mock.MockService{
				IsValidItemF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				IsValidUserF: func(ctx context.Context, id string) (bool, error) { return true, nil },
				GetItemF:     func(ctx context.Context, itemID, client string) (*models.Item, error) { return nil, nil },
			},
			wantStatus: http.StatusBadRequest,
			wantMsg:    NoUserFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerV1{meliService: tt.mock}

			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rr := httptest.NewRecorder()

			handler := timeoutMiddleware(1*time.Second, http.HandlerFunc(s.itemDetails))

			handler.ServeHTTP(rr, req)

			require.Equal(t, tt.wantStatus, rr.Code)

			var resp models.Response
			err := json.Unmarshal(rr.Body.Bytes(), &resp)
			require.NoError(t, err)
			require.Contains(t, resp.Message, tt.wantMsg)
		})
	}
}
