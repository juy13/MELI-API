package cache

import (
	"context"
	"itemmeli/constants"
	"itemmeli/models"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

var rdb *miniredis.Miniredis

func RedisCacheTest(t *testing.T) (*RedisCache, func()) {
	rdb = miniredis.RunT(t)
	if rdb == nil {
		t.Fatal("Could not run Redis server")
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     rdb.Addr(),
		Password: "",
		DB:       0,
	})
	return &RedisCache{
			client:                      redisClient,
			priceTTl:                    time.Second * 2,
			itemDetailsTTL:              time.Second * 2,
			customersRecommendationsTTL: time.Second * 2,
		}, func() {
			rdb.Close()
		}
}

func TestItemDetails(t *testing.T) {
	cache, cleanup := RedisCacheTest(t)
	defer cleanup()

	item := &models.Item{ID: constants.ItemID1, Title: "Test Item"}

	err := cache.SetItemDetails(context.Background(), constants.UserID1, constants.ItemID1, item)
	require.NoError(t, err)

	gotItem, err := cache.GetItemDetails(context.Background(), constants.UserID1, constants.ItemID1)
	require.NoError(t, err)
	require.NotNil(t, gotItem)
	require.Equal(t, "Test Item", gotItem.Title)
}

func TestItemDetailsExpired(t *testing.T) {
	cache, cleanup := RedisCacheTest(t)
	defer cleanup()

	item := &models.Item{ID: constants.ItemID1, Title: "Test Item"}

	err := cache.SetItemDetails(context.Background(), constants.UserID1, constants.ItemID1, item)
	require.NoError(t, err)
	rdb.FastForward(time.Second * 4) // miniredis doesn't have any timer inside, so we just moving the time

	gotItem, err := cache.GetItemDetails(context.Background(), constants.UserID1, constants.ItemID1)
	require.Nil(t, err)
	require.Nil(t, gotItem)
}

func TestItemPrice(t *testing.T) {
	cache, cleanup := RedisCacheTest(t)
	defer cleanup()
	itemPrice := 10.99

	price := &models.Price{Amount: itemPrice}

	err := cache.SetItemPrice(context.Background(), constants.UserID1, constants.ItemID1, price)
	require.NoError(t, err)

	gotItem, err := cache.GetItemPrice(context.Background(), constants.UserID1, constants.ItemID1)
	require.NoError(t, err)
	require.NotNil(t, gotItem)
	require.Equal(t, itemPrice, gotItem.Amount)
}

func TestItemPriceExpired(t *testing.T) {
	cache, cleanup := RedisCacheTest(t)
	defer cleanup()
	itemPrice := 10.99

	price := &models.Price{Amount: itemPrice}

	err := cache.SetItemPrice(context.Background(), constants.UserID1, constants.ItemID1, price)
	require.NoError(t, err)
	rdb.FastForward(time.Second * 4)

	gotItem, err := cache.GetItemPrice(context.Background(), constants.UserID1, constants.ItemID1)
	require.Nil(t, err)
	require.Nil(t, gotItem)
}

func TestCustomersRecommendations(t *testing.T) {
	cache, cleanup := RedisCacheTest(t)
	defer cleanup()
	items := []models.ItemShort{
		{ID: "item1", Title: "Item 1"},
		{ID: "item2", Title: "Item 2"},
	}

	err := cache.SetCustomersRecommendations(context.Background(), constants.UserID1, constants.ItemID1, constants.SellerID1, items)
	require.NoError(t, err)

	gotItem, err := cache.GetCustomersRecommendations(context.Background(), constants.UserID1, constants.ItemID1, constants.SellerID1)
	require.NoError(t, err)
	require.NotNil(t, gotItem)
	require.Equal(t, len(items), len(gotItem))

	for i := range items {
		require.Equal(t, items[i].ID, gotItem[i].ID)
		require.Equal(t, items[i].Title, gotItem[i].Title)
	}
}

func TestCustomersRecommendationsExpired(t *testing.T) {
	cache, cleanup := RedisCacheTest(t)
	defer cleanup()
	items := []models.ItemShort{
		{ID: "item1", Title: "Item 1"},
		{ID: "item2", Title: "Item 2"},
	}

	err := cache.SetCustomersRecommendations(context.Background(), constants.UserID1, constants.ItemID1, constants.SellerID1, items)
	require.NoError(t, err)
	rdb.FastForward(time.Second * 4)

	gotItem, err := cache.GetCustomersRecommendations(context.Background(), constants.UserID1, constants.SellerID1, constants.ItemID1)
	require.Nil(t, err)
	require.Nil(t, gotItem)
}
