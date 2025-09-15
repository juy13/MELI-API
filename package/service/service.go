package service

import (
	"context"
	"itemmeli/models"
)

type Service interface {
	GetItemDetails(ctx context.Context, itemID, clientID string) (*models.Item, error)
	GetItemRecommendations(ctx context.Context, itemID, sellerID, userID string) ([]models.ItemShort, error)

	// Checking if the user is valid or not
	IsValidUser(ctx context.Context, userID string) (bool, error)
	IsValidItem(ctx context.Context, itemID string) (bool, error)
}
