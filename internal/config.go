package config

import (
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Addr     string `json:"addr" `
	Password string `json:"password" `
	Login    string `json:"login" `
}

func CreateConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if strings.TrimSpace(configPath) == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Errorf("config file does not exist - `%s`", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
