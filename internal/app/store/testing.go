package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(
	t *testing.T,
	dbHost string,
	dbPort string,
	dbDatabase string,
	dbUser string,
	dbPassword string,
	dbSSLMode string,
) (*Store, func(...string)) {
	t.Helper()

	config, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	config.DBHost = dbHost
	config.DBPort = dbPort
	config.DBDatabase = dbDatabase
	config.DBUser = dbUser
	config.DBPassword = dbPassword
	config.DBSSLMode = dbSSLMode

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
