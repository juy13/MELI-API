package config

type Config interface {
	APIConfig
	CacheConfig
	MetricsConfig
}

type APIConfig interface {
	Port() int
	Host() string
}

type CacheConfig interface {
	CacheAddress() string
	CachePassword() string
	CacheDB() int
	MaxTweets2Keep() int
	TweetExpireTimeMinutes() int
	UserFeedExpireTimeMinutes() int
	TweetTimelineExpireTimeMinutes() int
	MaxTweetsTimelineItems() int
}

type MetricsConfig interface {
	MetricsServerHost() string
	MetricsServerPort() int
}
