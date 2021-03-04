package config

import (
	"fmt"
	"os"
	"path"

	uberConf "go.uber.org/config"
)

const (
	defaultConfDir = "conf/"
	defaultConfEnv = "local"
	DevEnv         = "dev"
	ProdEnv        = "prod"
)

type (
	Config struct {
		Transports Transports `yaml:"transports"`
		Logger     Logger     `yaml:"logger"`
	}

	Logger struct {
		Env string `yaml:"env"`
	}

	Transports struct {
		HTTP HTTPTransportConfig `yaml:"http"`
	}
	HTTPTransportConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
)

// NewConfig get config
func NewConfig() (*Config, error) {
	config := &Config{}

	// Get env
	confEnv := os.Getenv("CONF_ENV")
	if confEnv == "" {
		confEnv = defaultConfEnv
	}

	// From files
	file := fmt.Sprintf("%s.yml", confEnv)

	// Parse
	provider, err := uberConf.NewYAML(uberConf.File(path.Join(defaultConfDir, file)))
	if err != nil {
		return nil, err
	}

	// Populate
	err = provider.Get(uberConf.Root).Populate(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
