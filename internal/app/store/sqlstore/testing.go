package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDB(
	t *testing.T,
	dbHost string,
	dbPort string,
	dbDatabase string,
	dbUser string,
	dbPassword string,
	dbSSLMode string,
) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open(
		"postgres",
		"host="+dbHost+
			" port="+dbPort+
			" dbname="+dbDatabase+
			" user="+dbUser+
			" password="+dbPassword+
			" sslmode="+dbSSLMode,
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
