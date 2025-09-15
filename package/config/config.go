package config

import "time"

type Config interface {
	APIConfig
	MetricsConfig
	CacheConfig
	DatabaseConfig
}

type APIConfig interface {
	Port() int
	Host() string
	RequestTimeout() time.Duration
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

type DatabaseConfig interface {
	DBPath() string
	DBPath2() string // for recommendations
}
