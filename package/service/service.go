package service

import "context"

type Service interface {
	GetItemDetails(ctx context.Context, itemID string, clientID string) (interface{}, error)
	GetItemRecommendations(ctx context.Context, itemID string, clientID string) ([]interface{}, error)
}
