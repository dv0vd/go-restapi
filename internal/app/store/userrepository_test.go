package store_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/model"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store"
)

func init() {
	os.Setenv("APISERVER_CONFIG_PATH", "../../../configs/apiserver.toml")
}

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(
		t,
		dbHost,
		dbPort,
		dbDatabase,
		dbUser,
		dbPassword,
		dbSSLMode,
	)

	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(
		t,
		dbHost,
		dbPort,
		dbDatabase,
		dbUser,
		dbPassword,
		dbSSLMode,
	)

	defer teardown("users")

	email := "example@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
