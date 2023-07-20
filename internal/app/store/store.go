package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepositoty *UserRepositoty
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open(
		"postgres",
		"host="+s.config.DBHost+
			" port="+s.config.DBPort+
			" dbname="+s.config.DBDatabase+
			" user="+s.config.DBUser+
			" password="+s.config.DBPassword+
			" sslmode="+s.config.DBSSLMode,
	)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepositoty {
	if s.userRepositoty != nil {
		return s.userRepositoty
	}

	s.userRepositoty = &UserRepositoty{
		store: s,
	}

	return s.userRepositoty
}
