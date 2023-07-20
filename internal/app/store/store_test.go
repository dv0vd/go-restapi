package store_test

import (
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

var (
	dbHost     string
	dbPort     string
	dbDatabase string
	dbUser     string
	dbPassword string
	dbSSLMode  string
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../../.env"); err == nil {
		if dbHost = strings.ReplaceAll(os.Getenv("TEST_DB_HOST"), " ", ""); dbHost == "" {
			dbHost = "localhost"
		}

		if dbPort = strings.ReplaceAll(os.Getenv("TEST_DB_PORT"), " ", ""); dbPort == "" {
			dbPort = "5432"
		}

		if dbDatabase = strings.ReplaceAll(os.Getenv("TEST_DB_DATABASE"), " ", ""); dbDatabase == "" {
			dbDatabase = "database_test"
		}

		if dbUser = strings.ReplaceAll(os.Getenv("TEST_DB_USERNAME"), " ", ""); dbUser == "" {
			dbUser = "postgres"
		}

		if dbPassword = strings.ReplaceAll(os.Getenv("TEST_DB_PASSWORD"), " ", ""); dbPassword == "" {
			dbPassword = "postgres"
		}

		if dbSSLMode = strings.ReplaceAll(os.Getenv("TEST_DB_SSL_MODE"), " ", ""); dbSSLMode == "" {
			dbSSLMode = "require"
		}
	}

	os.Exit(m.Run())
}
