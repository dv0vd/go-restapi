package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepositoty *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepositoty != nil {
		return s.userRepositoty
	}

	s.userRepositoty = &UserRepository{
		store: s,
	}

	return s.userRepositoty
}
