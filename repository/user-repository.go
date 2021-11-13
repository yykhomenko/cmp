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
	return &database{
		data: make(map[string]entity.User),
	}
}

func (db *database) Get(login string) entity.User {
	return db.data[login]
}

func (db *database) Set(user entity.User) {
	db.data[user.Login] = user
}
