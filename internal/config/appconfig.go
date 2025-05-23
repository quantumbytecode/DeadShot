package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfigs() (*AppConfig, error) {
	basePath, _ := os.Getwd()

	configPath := filepath.Join(basePath, "internal", "config", "appconfig.yaml")

	f, e := os.ReadFile(configPath)

	if e != nil {
		log.Error("Could not load configurations")
		log.Error(e)
		return nil, e
	}

	var config AppConfig

	yaml.Unmarshal(f, &config)

	return &config, nil
}
