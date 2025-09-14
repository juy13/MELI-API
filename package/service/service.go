package service

import (
	"context"
	"itemmeli/package/models"
)

type Service interface {
	GetItemDetails(ctx context.Context, itemID string, clientID string) (*models.Item, error)
	GetItemRecommendations(ctx context.Context, itemID string, sellerID, clientID string) ([]models.ItemShort, error)
}
