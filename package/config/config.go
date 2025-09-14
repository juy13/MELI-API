package config

import "time"

type Config interface {
	APIConfig
	MetricsConfig
	CacheConfig
}

type APIConfig interface {
	Port() int
	Host() string
}

type MetricsConfig interface {
	MetricsServerHost() string
	MetricsServerPort() int
}

type CacheConfig interface {
	CacheAddress() string
	CachePassword() string
	CacheDB() int
	PriceTTL() time.Duration
	ItemDetailsTTL() time.Duration
	CustomersRecommendationsTTL() time.Duration
}
