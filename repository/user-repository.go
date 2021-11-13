package repository

import "github.com/yykhomenko/cmp/entity"

type UserRepository interface {
	Get(login string) entity.User
	Set(entity.User)
}

type database struct {
	data map[string]entity.User
}

func NewUserRepository() UserRepository {
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

func (db *database) Set(user entity.User) {
	db.data[user.Login] = user
}
