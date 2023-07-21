package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(
		config.DBHost,
		config.DBPort,
		config.DBDatabase,
		config.DBUser,
		config.DBPassword,
		config.DBSSLMode,
	)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := newServer(store, sessionStore)

	return http.ListenAndServe(config.AppUrl+":"+config.AppPort, s)
}

func newDB(
	dbHost string,
	dbPort string,
	dbDatabase string,
	dbUser string,
	dbPassword string,
	dbSSLMode string,
) (*sql.DB, error) {
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
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
