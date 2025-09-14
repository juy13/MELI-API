package config

import (
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	// api
	API APIStruct `yaml:"api,omitempty"`

	// metrics
	MetricsServer MetricsConfigStruct `yaml:"metrics,omitempty"`
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
