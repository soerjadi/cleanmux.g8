package config

import (
	"os"

	"main/internal/pkg/util"

	"gopkg.in/gcfg.v1"
)

var configFilePaths = map[string]string{
	"PRODUCTION":  "$production_config$",
	"DEVELOPMENT": "../../files/config.ini",
}

func Init() (*Config, error) {
	cfg = &Config{}

	configFilePath := configFilePaths[util.GetENV()]

	config, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = gcfg.ReadStringInto(cfg, string(config))
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// GetConfig returns config object
func GetConfig() *Config {
	return cfg
}
