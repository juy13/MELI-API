package database

import (
	"context"
	"itemmeli/package/models"
)

type Database interface {
	GetItemDetails(ctx context.Context, itemID string) (*models.Item, error)
	GetItemRecommendations(ctx context.Context, sellerID, itemID string) ([]models.ItemShort, error)
}
