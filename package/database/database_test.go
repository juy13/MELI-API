package database

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

type LocalConfig struct {
	Path  string
	Path2 string
}

func (l *LocalConfig) DBPath() string {
	return l.Path
}

func (l *LocalConfig) DBPath2() string {
	return l.Path2
}

var localConfig = &LocalConfig{
	Path:  filepath.Join("testdata", "test_items.json"),
	Path2: filepath.Join("testdata", "test_recommendations.json"),
}

func TestJSONDatabase_LoadAndGetItem(t *testing.T) {
	db, err := NewJSONDatabase(localConfig)
	require.NoError(t, err)
	require.NotNil(t, db)

	t.Run("Get existing item", func(t *testing.T) {
		item, err := db.GetItemDetails(context.Background(), "item1")
		require.NoError(t, err)
		require.NotNil(t, item)
		require.Equal(t, "Red Shoes", item.Title)
		require.Equal(t, 79.99, item.Price.Amount)
		require.True(t, item.Available)
	})

	t.Run("Get missing item", func(t *testing.T) {
		item, err := db.GetItemDetails(context.Background(), "does-not-exist")
		require.Error(t, err)
		require.Nil(t, item)
	})
}

func TestJSONDatabase_Recommendations(t *testing.T) {
	db, err := NewJSONDatabase(localConfig)
	require.NoError(t, err)

	recs, err := db.GetItemRecommendations(context.Background(), "seller1", "item1")
	require.NoError(t, err)

	require.Len(t, recs, 1)

	first := recs[0]
	require.Equal(t, "item2", first.ID)
	require.Equal(t, "Comfortable Leather Shoes", first.Title)
	require.Equal(t, "ARS", first.Price.CurrencyID)
	require.Equal(t, "tomorrow", first.Shipping.Time)
}
