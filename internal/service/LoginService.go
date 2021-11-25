package service

import "github.com/yykhomenko/cmp/internal/repository"

type LoginService interface {
	LoginUser(login string, password string) bool
}

type loginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(userRepository repository.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (ls *loginService) LoginUser(login string, password string) bool {
	user := ls.userRepository.Get(login)
	return user.Login == login && user.Password == password
}
