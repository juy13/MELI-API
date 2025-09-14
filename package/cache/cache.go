package cache

import (
	"context"
	"itemmeli/models"
)

type Cache interface {
	GetItemDetails(ctx context.Context, userID, itemID string) (*models.Item, error)
	SetItemDetails(ctx context.Context, userID, itemID string, itemDetail *models.Item) error

	// we assume that the price can change all the time and let's keep it in cache separately
	// also price can depend on the user and in future on his device, timestamp and etc
	GetItemPrice(ctx context.Context, userID, itemID string) (*models.Price, error)
	SetItemPrice(ctx context.Context, userID, itemID string, price *models.Price) error

	GetCustomersRecommendations(ctx context.Context, userID, itemID string) ([]models.ItemShort, error)
	SetCustomersRecommendations(ctx context.Context, userID, itemID string, recommendations []models.ItemShort) error
}
