package mock

import (
	"context"
	"itemmeli/models"
)

type MockService struct {
	IsValidItemF            func(ctx context.Context, itemID string) (bool, error)
	IsValidUserF            func(ctx context.Context, userID string) (bool, error)
	GetItemF                func(ctx context.Context, itemID, client string) (*models.Item, error)
	GetItemRecommendationsF func(ctx context.Context, itemID, sellerID, userID string) ([]models.ItemShort, error)
}

func (m *MockService) IsValidItem(ctx context.Context, itemID string) (bool, error) {
	return m.IsValidItemF(ctx, itemID)
}
func (m *MockService) IsValidUser(ctx context.Context, userID string) (bool, error) {
	return m.IsValidUserF(ctx, userID)
}
func (m *MockService) GetItemDetails(ctx context.Context, itemID, client string) (*models.Item, error) {
	return m.GetItemF(ctx, itemID, client)
}
func (m *MockService) GetItemRecommendations(ctx context.Context, itemID, sellerID, userID string) ([]models.ItemShort, error) {
	return m.GetItemRecommendationsF(ctx, itemID, sellerID, userID)
}
