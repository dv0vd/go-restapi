package apiserver

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	AppUrl      string `toml:"app_url"`
	AppPort     string `toml:"app_port"`
	AppLogLevel string `toml:"app_log_level"`
}

func NewConfig() (*Config, error) {
	var config Config

	_, err := toml.DecodeFile(os.Getenv("APISERVER_CONFIG_PATH"), &config)
	if err != nil {
		return nil, err
	}

	var (
		appUrl      string
		appPort     string
		appLogLevel string
	)
	if err := godotenv.Load(); err == nil {
		if appUrl = strings.ReplaceAll(os.Getenv("APP_URL"), " ", ""); appUrl != "" {
			config.AppUrl = appUrl
		}

		if appPort = strings.ReplaceAll(os.Getenv("APP_PORT"), " ", ""); appPort != "" {
			config.AppPort = appPort
		}

		if appLogLevel = strings.ReplaceAll(os.Getenv("APP_LOG_LEVEL"), " ", ""); appLogLevel != "" {
			config.AppLogLevel = appLogLevel
		}
	}

	return &config, nil
}
