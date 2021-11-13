package service

import "github.com/yykhomenko/cmp/repository"

type LoginService interface {
	LoginUser(login string, password string) bool
}

type loginService struct {
	userRepo repository.UserRepository
}

func NewLoginService(userRepo repository.UserRepository) LoginService {
	return &loginService{
		userRepo: userRepo,
	}
}

func (ls *loginService) LoginUser(login string, password string) bool {
	user := ls.userRepo.Get(login)
	return user.Password == password
}
