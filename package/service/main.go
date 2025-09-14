package service

import (
	"context"
	"itemmeli/package/cache"
	"itemmeli/package/database"
	"itemmeli/package/models"
)

type APIService struct {
	cache cache.Cache
	db    database.Database

	// priceService
	// Shipping Service -- to know price for shipping and time to come
}

func NewService(cache cache.Cache, db database.Database) Service {
	return &APIService{
		cache: cache,
		db:    db,
	}
}

func (api *APIService) GetItemDetails(ctx context.Context, itemID string, clientID string) (*models.Item, error) {
	cachedItem, err := api.cache.GetItemDetails(ctx, clientID, itemID)
	if err == nil && cachedItem != nil {
		return cachedItem, nil
	}

	item, err := api.db.GetItemDetails(ctx, itemID)
	if err != nil {
		return nil, err
	}

	// updating price in a case if there are any discounts for user, for example
	item.Price.Amount = priceOracle(item.Price.Amount)

	_ = api.cache.SetItemDetails(ctx, clientID, itemID, item)

	return item, nil
}

func (api *APIService) GetItemRecommendations(ctx context.Context, itemID string, sellerID, clientID string) ([]models.ItemShort, error) {
	cachedRecs, err := api.cache.GetCustomersRecommendations(ctx, clientID, itemID)
	if err == nil && cachedRecs != nil {
		return cachedRecs, nil
	}

	recs, err := api.db.GetItemRecommendations(ctx, sellerID, itemID)
	if err != nil {
		return nil, err
	}

	// we are ignoring redis errors cz if it fails -- nothing wrong, we should have a monitoring for this case.
	_ = api.cache.SetCustomersRecommendations(ctx, clientID, itemID, recs)

	return recs, nil
}

func priceOracle(basePrice float64) float64 {
	return basePrice * 1.05 // idk, it could be another service also, I've sent a question wtd with it
}
