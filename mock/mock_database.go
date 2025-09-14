package mock

import (
	"context"
	"itemmeli/package/models"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) GetItemDetails(ctx context.Context, itemID string) (*models.Item, error) {
	args := m.Called(ctx, itemID)
	if item, ok := args.Get(0).(*models.Item); ok {
		return item, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockDB) GetItemRecommendations(ctx context.Context, sellerID, itemID string) ([]models.ItemShort, error) {
	args := m.Called(ctx, sellerID, itemID)
	if recs, ok := args.Get(0).([]models.ItemShort); ok {
		return recs, args.Error(1)
	}
	return nil, args.Error(1)
}
