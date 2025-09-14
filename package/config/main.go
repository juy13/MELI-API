package config

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	// api
	API APIStruct `yaml:"api,omitempty"`

	// metrics
	MetricsServer MetricsConfigStruct `yaml:"metrics,omitempty"`

	// cache redis
	Redis CacheConfigStruct `yaml:"cache"`

	// database
	Database DatabaseConfigStruct `yaml:"database,omitempty"`
}

type DatabaseConfigStruct struct {
	Path  string `yaml:"path"`
	Path2 string `yaml:"path2"`
}

type CacheConfigStruct struct {
	CacheAddress   string        `yaml:"address"`
	CachePassword  string        `yaml:"password"`
	CacheDB        int           `yaml:"db"`
	PriceTTL       time.Duration `yaml:"price_ttl"`
	ItemDetailsTTL time.Duration `yaml:"item_details_ttl"`
	CustomersTTL   time.Duration `yaml:"customers_recommendations_ttl"`
}

type APIStruct struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type MetricsConfigStruct struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewYamlConfig(configFilePath string) (*YamlConfig, error) {
	var (
		err  error
		file *os.File
		data []byte
	)
	dc := &YamlConfig{}

	if file, err = os.Open(configFilePath); err != nil {
		log.Fatal().Msg("can't open config file")
	}
	defer func() {
		_ = file.Close()
	}() // TODO check closings

	if data, err = io.ReadAll(file); err != nil {
		log.Fatal().Msg("can't read config file")
	}
	if err = yaml.Unmarshal(data, dc); err != nil {
		log.Fatal().Msg("can't unmarshal config file")
	}
	return dc, nil
}

///////////////////////////////////
//	API
///////////////////////////////////

func (c YamlConfig) Port() int {
	return c.API.Port
}

func (c *YamlConfig) Host() string {
	return c.API.Host
}

///////////////////////////////////
//	Metrics Config
///////////////////////////////////

func (c *YamlConfig) MetricsServerHost() string {
	return c.MetricsServer.Host
}
func (c *YamlConfig) MetricsServerPort() int {
	return c.MetricsServer.Port
}

///////////////////////////////////
//	Redis Config
///////////////////////////////////

func (c *YamlConfig) CacheAddress() string {
	return c.Redis.CacheAddress
}
func (c *YamlConfig) CachePassword() string {
	return c.Redis.CachePassword
}
func (c *YamlConfig) CacheDB() int {
	return c.Redis.CacheDB
}
func (c *YamlConfig) PriceTTL() time.Duration {
	return c.Redis.PriceTTL
}
func (c *YamlConfig) ItemDetailsTTL() time.Duration {
	return c.Redis.ItemDetailsTTL
}
func (c *YamlConfig) CustomersRecommendationsTTL() time.Duration {
	return c.Redis.CustomersTTL
}

///////////////////////////////////
//	Database Config
///////////////////////////////////

func (c *YamlConfig) DBPath() string {
	return c.Database.Path
}

func (c *YamlConfig) DBPath2() string {
	return c.Database.Path2
}
