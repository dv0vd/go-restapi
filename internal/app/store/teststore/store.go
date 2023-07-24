package teststore

import (
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/model"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store"
)

type Store struct {
	UserRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.UserRepository
}
