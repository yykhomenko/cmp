package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yykhomenko/cmp/entity"
	"github.com/yykhomenko/cmp/repository"
)

func TestLoginUser(t *testing.T) {
	userRepository := repository.NewUserRepository()
	loginService := NewLoginService(userRepository)

	testUser := entity.User{
		Login:    "yykhomenko",
		Password: "123",
	}
	userRepository.Set(testUser)
	isLogin := loginService.LoginUser(testUser.Login, testUser.Password)

	assert.True(t, isLogin)
}
