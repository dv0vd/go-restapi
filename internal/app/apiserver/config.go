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
	DBHost      string `toml:"db_host"`
	DBPort      string `toml:"db_port"`
	DBDatabase  string `toml:"db_database"`
	DBUser      string `toml:"db_user"`
	DBPassword  string `toml:"db_password"`
	DBSSLMode   string `toml:"db_ssl_mode"`
	SessionKey  string `toml:"session_key"`
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
		dbHost      string
		dbPort      string
		dbDatabase  string
		dbUser      string
		dbPassword  string
		dbSSLMode   string
		sessionKey  string
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

		if dbHost = strings.ReplaceAll(os.Getenv("DB_HOST"), " ", ""); dbHost != "" {
			config.DBHost = dbHost
		}

		if dbPort = strings.ReplaceAll(os.Getenv("DB_PORT"), " ", ""); dbPort != "" {
			config.DBPort = dbPort
		}

		if dbDatabase = strings.ReplaceAll(os.Getenv("DB_DATABASE"), " ", ""); dbDatabase != "" {
			config.DBDatabase = dbDatabase
		}

		if dbUser = strings.ReplaceAll(os.Getenv("DB_USERNAME"), " ", ""); dbUser != "" {
			config.DBUser = dbUser
		}

		if dbPassword = strings.ReplaceAll(os.Getenv("DB_PASSWORD"), " ", ""); dbPassword != "" {
			config.DBPassword = dbPassword
		}

		if dbSSLMode = strings.ReplaceAll(os.Getenv("DB_SSL_MODE"), " ", ""); dbSSLMode != "" {
			config.DBSSLMode = dbSSLMode
		}

		if sessionKey = strings.ReplaceAll(os.Getenv("SESSION_KEY"), " ", ""); sessionKey != "" {
			config.SessionKey = sessionKey
		}
	}

	return &config, nil
}
