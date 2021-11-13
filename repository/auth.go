package repository

import "github.com/yykhomenko/cmp/entity"

type AuthRepository interface {
	Get(login string) entity.User
}

type database struct {
	data map[string]entity.User
}

func NewAuthRepository() AuthRepository {
	testUser := entity.User{
		Login:    "yykhomenko",
		Password: "123",
	}
	data := make(map[string]entity.User)
	data[testUser.Login] = testUser
	return &database{
		data: data,
	}
}

func (db *database) Get(login string) entity.User {
	return db.data[login]
}
