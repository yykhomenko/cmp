package repository

import "github.com/yykhomenko/cmp/internal/entity"

type UserRepository interface {
	Get(login string) entity.User
	Set(entity.User)
}

type userRepository struct {
	data map[string]entity.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		data: make(map[string]entity.User),
	}
}

func (ur *userRepository) Get(login string) entity.User {
	return ur.data[login]
}

func (ur *userRepository) Set(user entity.User) {
	ur.data[user.Login] = user
}
