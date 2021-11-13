package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	login := "yykhomenko"
	repo := NewAuthRepository()
	user := repo.Get(login)
	assert.Equal(t, login, user.Login)
}
