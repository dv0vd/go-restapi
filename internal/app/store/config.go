package store

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string `toml:"db_host"`
	DBPort     string `toml:"db_port"`
	DBDatabase string `toml:"db_database"`
	DBUser     string `toml:"db_user"`
	DBPassword string `toml:"db_password"`
	DBSSLMode  string `toml:"db_ssl_mode"`
}

func NewConfig() (*Config, error) {
	var config Config

	_, err := toml.DecodeFile(os.Getenv("APISERVER_CONFIG_PATH"), &config)
	if err != nil {
		return nil, err
	}

	var (
		dbHost     string
		dbPort     string
		dbDatabase string
		dbUser     string
		dbPassword string
		dbSSLMode  string
	)

	if err := godotenv.Load(); err == nil {
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
	}

	return &config, nil
}
