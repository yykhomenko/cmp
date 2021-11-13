package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yykhomenko/cmp/entity"
)

func TestSetGet(t *testing.T) {
	testUser := entity.User{
		Login:    "yykhomenko",
		Password: "123",
	}

	userRepository := NewUserRepository()
	userRepository.Set(testUser)
	actualUser := userRepository.Get(testUser.Login)

	assert.Equal(t, testUser.Login, actualUser.Login)
	assert.Equal(t, testUser.Password, actualUser.Password)
}
