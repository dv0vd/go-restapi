package sqlstore_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/model"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store/sqlstore"
)

func init() {
	os.Setenv("APISERVER_CONFIG_PATH", "../../../../configs/apiserver.toml")
}

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(
		t,
		dbHost,
		dbPort,
		dbDatabase,
		dbUser,
		dbPassword,
		dbSSLMode,
	)

	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(
		t,
		dbHost,
		dbPort,
		dbDatabase,
		dbUser,
		dbPassword,
		dbSSLMode,
	)

	defer teardown("users")

	s := sqlstore.New(db)

	email := "example@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
