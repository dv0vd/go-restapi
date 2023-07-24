package teststore_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/model"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store"
	"gitlab.qsoft.ru/grade/v.davydov_first_rest_api/internal/app/store/teststore"
)

func init() {
	os.Setenv("APISERVER_CONFIG_PATH", "../../../../configs/apiserver.toml")
}

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	u1 := model.TestUser(t)
	s.User().Create(u1)

	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	email := "example@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u1 := model.TestUser(t)
	u1.Email = email
	s.User().Create(u1)

	u2, err := s.User().FindByEmail(u1.Email)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
