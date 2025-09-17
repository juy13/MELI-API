package mock

import (
	"context"
	"itemmeli/models"

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
func (m *MockDB) GetItemRecommendations(ctx context.Context, userID, itemID, sellerID string) ([]models.ItemShort, error) {
	args := m.Called(ctx, userID, itemID, sellerID)
	if recs, ok := args.Get(0).([]models.ItemShort); ok {
		return recs, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDB) GetUser(ctx context.Context, userID string) (*models.User, error) {
	args := m.Called(ctx, userID)
	if recs, ok := args.Get(0).(models.User); ok {
		return &recs, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDB) GetItem(ctx context.Context, itemID string) (*models.ItemShort, error) {
	args := m.Called(ctx, itemID)
	if recs, ok := args.Get(0).(models.ItemShort); ok {
		return &recs, args.Error(1)
	}
	return nil, args.Error(1)
}
