package database

import (
	"context"
	"itemmeli/models"
)

type Database interface {
	GetItemDetails(ctx context.Context, itemID string) (*models.Item, error)
	GetItemRecommendations(ctx context.Context, userID, itemID, sellerID string) ([]models.ItemShort, error)

	// Items and users
	GetUser(ctx context.Context, userID string) (*models.User, error)
	GetItem(ctx context.Context, itemID string) (*models.ItemShort, error)
}
