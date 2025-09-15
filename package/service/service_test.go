package service

import (
	"context"
	"errors"
	"itemmeli/mock"
	"itemmeli/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetItemDetails_CacheHit(t *testing.T) {
	ctx := context.Background()
	cache := new(mock.MockCache)
	db := new(mock.MockDB)
	api := NewService(cache, db)

	expected := &models.Item{ID: "item1", Title: "Test Item"}

	cache.On("GetItemDetails", ctx, "client1", "item1").Return(expected, nil)

	item, err := api.GetItemDetails(ctx, "item1", "client1")
	require.NoError(t, err)
	require.Equal(t, expected, item)

	cache.AssertExpectations(t)
	db.AssertExpectations(t) // should not be called
}

func TestGetItemDetails_CacheMiss_DBHit(t *testing.T) {
	ctx := context.Background()
	cache := new(mock.MockCache)
	db := new(mock.MockDB)
	api := NewService(cache, db)

	expected := &models.Item{ID: "item1", Title: "From DB"}

	cache.On("GetItemDetails", ctx, "client1", "item1").Return(nil, errors.New("cache miss"))
	db.On("GetItemDetails", ctx, "item1").Return(expected, nil)
	cache.On("SetItemDetails", ctx, "client1", "item1", expected).Return(nil)

	item, err := api.GetItemDetails(ctx, "item1", "client1")
	require.NoError(t, err)
	require.Equal(t, expected, item)

	cache.AssertExpectations(t)
	db.AssertExpectations(t)
}

func TestGetItemRecommendations_CacheHit(t *testing.T) {
	ctx := context.Background()
	cache := new(mock.MockCache)
	db := new(mock.MockDB)
	api := NewService(cache, db)

	expected := []models.ItemShort{{ID: "rec1", Title: "Cached Rec"}}

	cache.On("GetCustomersRecommendations", ctx, "client1", "item1").Return(expected, nil)

	recs, err := api.GetItemRecommendations(ctx, "item1", "seller1", "client1")
	require.NoError(t, err)
	require.Equal(t, expected, recs)

	cache.AssertExpectations(t)
	db.AssertExpectations(t) // DB should not be called
}

func TestGetItemRecommendations_CacheMiss_DBHit(t *testing.T) {
	ctx := context.Background()
	cache := new(mock.MockCache)
	db := new(mock.MockDB)
	api := NewService(cache, db)

	expected := []models.ItemShort{{ID: "rec1", Title: "From DB"}}

	cache.On("GetCustomersRecommendations", ctx, "client1", "item1").Return(nil, errors.New("cache miss"))
	db.On("GetItemRecommendations", ctx, "seller1", "item1").Return(expected, nil)
	cache.On("SetCustomersRecommendations", ctx, "client1", "item1", expected).Return(nil)

	recs, err := api.GetItemRecommendations(ctx, "item1", "seller1", "client1")
	require.NoError(t, err)
	require.Equal(t, expected, recs)

	cache.AssertExpectations(t)
	db.AssertExpectations(t)
}
