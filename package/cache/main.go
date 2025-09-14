package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"itemmeli/package/config"
	"itemmeli/package/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client

	priceTTl                    time.Duration
	itemDetailsTTL              time.Duration
	customersRecommendationsTTL time.Duration
}

func NewRedisCache(config config.Config) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.CacheAddress(),
		Password: config.CachePassword(),
		DB:       config.CacheDB(),
	})

	return &RedisCache{
		client:                      rdb,
		priceTTl:                    config.PriceTTL(),
		itemDetailsTTL:              config.ItemDetailsTTL(),
		customersRecommendationsTTL: config.CustomersRecommendationsTTL(),
	}
}

func itemDetailsKey(userID, itemID string) string {
	return fmt.Sprintf("item:details:%s:%s", userID, itemID)
}

func itemPriceKey(userID, itemID string) string {
	return fmt.Sprintf("item:price:%s:%s", userID, itemID)
}

func recsKey(userID string) string {
	return fmt.Sprintf("recs:%s", userID)
}

func (c *RedisCache) GetItemDetails(ctx context.Context, userID, itemID string) (*models.Item, error) {
	val, err := c.client.Get(ctx, itemDetailsKey(userID, itemID)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var item models.Item
	if err := json.Unmarshal([]byte(val), &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *RedisCache) SetItemDetails(ctx context.Context, userID, itemID string, itemDetail *models.Item) error {
	data, err := json.Marshal(itemDetail)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, itemDetailsKey(userID, itemID), data, c.itemDetailsTTL).Err()
}

func (c *RedisCache) GetItemPrice(ctx context.Context, userID, itemID string) (*models.Price, error) {
	val, err := c.client.Get(ctx, itemPriceKey(userID, itemID)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var price models.Price
	if err := json.Unmarshal([]byte(val), &price); err != nil {
		return nil, err
	}
	return &price, nil
}

func (c *RedisCache) SetItemPrice(ctx context.Context, userID, itemID string, price *models.Price) error {
	data, err := json.Marshal(price)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, itemPriceKey(userID, itemID), data, c.priceTTl).Err()
}

func (c *RedisCache) GetCustomersRecommendations(ctx context.Context, userID, itemID string) ([]models.ItemShort, error) {
	val, err := c.client.Get(ctx, recsKey(userID)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var recs []models.ItemShort
	if err := json.Unmarshal([]byte(val), &recs); err != nil {
		return nil, err
	}
	return recs, nil
}

func (c *RedisCache) SetCustomersRecommendations(ctx context.Context, userID, itemID string, recommendations []models.ItemShort) error {
	data, err := json.Marshal(recommendations)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, recsKey(userID), data, c.customersRecommendationsTTL).Err()
}
