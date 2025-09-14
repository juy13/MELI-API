package mock

import (
	"context"
	"itemmeli/package/models"

	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (m *MockCache) GetItemDetails(ctx context.Context, userID, itemID string) (*models.Item, error) {
	args := m.Called(ctx, userID, itemID)
	if item, ok := args.Get(0).(*models.Item); ok {
		return item, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockCache) SetItemDetails(ctx context.Context, userID, itemID string, item *models.Item) error {
	return m.Called(ctx, userID, itemID, item).Error(0)
}

func (m *MockCache) GetCustomersRecommendations(ctx context.Context, userID, itemID string) ([]models.ItemShort, error) {
	args := m.Called(ctx, userID, itemID)
	if recs, ok := args.Get(0).([]models.ItemShort); ok {
		return recs, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockCache) SetCustomersRecommendations(ctx context.Context, userID, itemID string, recs []models.ItemShort) error {
	return m.Called(ctx, userID, itemID, recs).Error(0)
}

func (m *MockCache) GetItemPrice(ctx context.Context, userID, itemID string) (*models.Price, error) {
	args := m.Called(ctx, userID, itemID)
	if recs, ok := args.Get(0).(*models.Price); ok {
		return recs, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockCache) SetItemPrice(ctx context.Context, userID, itemID string, price *models.Price) error {
	return m.Called(ctx, userID, itemID, price).Error(0)
}
