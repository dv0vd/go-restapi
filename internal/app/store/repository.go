package store

import "gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
