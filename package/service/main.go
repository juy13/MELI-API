package service

import (
	"context"
	"itemmeli/package/cache"
)

type APIService struct {
}

func NewService(cache cache.Cache) Service {
	var s APIService
	return &s
}

func (api *APIService) GetItemDetails(ctx context.Context, itemID string, clientID string) (interface{}, error) {
	panic("implement me")
}
func (api *APIService) GetItemRecommendations(ctx context.Context, itemID string, clientID string) ([]interface{}, error) {
	panic("implement me")
}
