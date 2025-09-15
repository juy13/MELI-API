package database

import (
	"context"
	"encoding/json"
	"fmt"
	"itemmeli/models"
	"itemmeli/package/config"
	"os"
)

type JSONDatabase struct {
	items           map[string]models.Item
	recommendations map[string]map[string][]models.ItemShort
	users           map[string]models.User
}

func NewJSONDatabase(config config.DatabaseConfig) (*JSONDatabase, error) {
	f, err := os.Open(config.DBPath())
	if err != nil {
		return nil, fmt.Errorf("open json: %w", err)
	}
	defer f.Close()

	var items []models.Item
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	itemMap := make(map[string]models.Item)
	for _, it := range items {
		itemMap[it.ID] = it
	}

	fRec, err := os.Open(config.DBPath2())
	if err != nil {
		return nil, fmt.Errorf("open json: %w", err)
	}
	defer fRec.Close()

	recommendations := make(map[string]map[string][]models.ItemShort)
	if err := json.NewDecoder(fRec).Decode(&recommendations); err != nil {
		return nil, fmt.Errorf("decode json: %w", err)
	}

	return &JSONDatabase{items: itemMap, recommendations: recommendations}, nil
}

func (db *JSONDatabase) GetItemDetails(ctx context.Context, itemID string) (*models.Item, error) {
	it, ok := db.items[itemID]
	if !ok {
		return nil, fmt.Errorf("item not found: %s", itemID)
	}
	return &it, nil
}

func (db *JSONDatabase) GetItemRecommendations(ctx context.Context, sellerID, itemID string) ([]models.ItemShort, error) {
	return db.recommendations[sellerID][itemID], nil
}

func (db *JSONDatabase) GetUser(ctx context.Context, userID string) (*models.User, error) {
	if user, ok := db.users[userID]; ok {
		return &user, nil
	}
	return nil, fmt.Errorf("not implemented")
}

func (db *JSONDatabase) GetItem(ctx context.Context, itemID string) (*models.ItemShort, error) {
	if it, ok := db.items[itemID]; ok {
		return &models.ItemShort{
			ID:       it.ID,
			Title:    it.Title,
			Price:    it.Price,
			Shipping: it.Shipping,
		}, nil
	}
	return nil, fmt.Errorf("item not found: %s", itemID)
}
