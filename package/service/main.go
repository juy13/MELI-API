package service

import (
	"context"
	"fmt"
	"itemmeli/models"
	"itemmeli/package/cache"
	"itemmeli/package/database"
	"regexp"
)

type APIService struct {
	cache cache.Cache
	db    database.Database

	// priceService
	// Shipping Service -- to know price for shipping and time to come

	reCheckEntry *regexp.Regexp
}

func NewService(cache cache.Cache, db database.Database) Service {
	re, err := regexp.Compile(`^[0-9-]*$`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return nil
	}
	return &APIService{
		cache:        cache,
		db:           db,
		reCheckEntry: re,
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

func (api *APIService) GetItemRecommendations(ctx context.Context, itemID string, sellerID, userID string) ([]models.ItemShort, error) {
	cachedRecs, err := api.cache.GetCustomersRecommendations(ctx, userID, itemID)
	if err == nil && cachedRecs != nil {
		return cachedRecs, nil
	}

	recs, err := api.db.GetItemRecommendations(ctx, sellerID, itemID)
	if err != nil {
		return nil, err
	}

	// we are ignoring redis errors cz if it fails -- nothing wrong, we should have a monitoring for this case.
	_ = api.cache.SetCustomersRecommendations(ctx, userID, itemID, recs)

	return recs, nil
}

func (api *APIService) IsValidItem(ctx context.Context, itemID string) (bool, error) {
	if !api.CheckEntry(itemID) {
		return false, fmt.Errorf("item not matching pattern")
	}
	exists, err := api.cache.CheckItem(ctx, itemID)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}

	_, err = api.db.GetItem(ctx, itemID)
	if err != nil {
		return false, err
	}

	// TODO set item in cache

	return true, nil
}

func (api *APIService) IsValidUser(ctx context.Context, userID string) (bool, error) {
	if !api.CheckEntry(userID) {
		return false, fmt.Errorf("user not matching pattern")
	}
	exists, err := api.cache.CheckUser(ctx, userID)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}

	_, err = api.db.GetItem(ctx, userID)
	if err != nil {
		return false, err
	}

	// TODO set user in cache
	return true, nil
}

func (api APIService) CheckEntry(entry string) bool {
	return api.reCheckEntry.MatchString(entry)
}

func priceOracle(basePrice float64) float64 {
	return basePrice * 1.05 // idk, it could be another service also, I've sent a question wtd with it
}
