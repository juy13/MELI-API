package mock

import (
	"context"
	"itemmeli/package/models"
)

type MockCache struct {
	itemDetails map[string]*models.Item       // key: userID|itemID
	itemPrices  map[string]*models.Price      // key: userID|itemID
	recs        map[string][]models.ItemShort // key: userID
}

func NewMockCache() *MockCache {
	return &MockCache{
		itemDetails: make(map[string]*models.Item),
		itemPrices:  make(map[string]*models.Price),
		recs:        make(map[string][]models.ItemShort),
	}
}

// build key helper
func makeKey(userID, itemID string) string {
	return userID + "|" + itemID
}

func (c *MockCache) GetItemDetails(ctx context.Context, userID, itemID string) (*models.Item, error) {
	key := makeKey(userID, itemID)
	if val, ok := c.itemDetails[key]; ok {
		return val, nil
	}
	return nil, nil // not found
}

func (c *MockCache) SetItemDetails(ctx context.Context, userID, itemID string, itemDetail *models.Item) error {
	key := makeKey(userID, itemID)
	c.itemDetails[key] = itemDetail
	return nil
}

func (c *MockCache) GetItemPrice(ctx context.Context, userID, itemID string) (*models.Price, error) {
	key := makeKey(userID, itemID)
	if val, ok := c.itemPrices[key]; ok {
		return val, nil
	}
	return nil, nil
}

func (c *MockCache) SetItemPrice(ctx context.Context, userID, itemID string, price *models.Price) error {
	key := makeKey(userID, itemID)
	c.itemPrices[key] = price
	return nil
}

func (c *MockCache) GetCustomersRecommendations(ctx context.Context, userID string) ([]models.ItemShort, error) {
	if val, ok := c.recs[userID]; ok {
		return val, nil
	}
	return nil, nil
}

func (c *MockCache) SetCustomersRecommendations(ctx context.Context, userID string, recommendations []models.ItemShort) {
	c.recs[userID] = recommendations
}
