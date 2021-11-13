package repository

import "github.com/yykhomenko/cmp/entity"

type AuthRepository interface {
	Get(name string) entity.User
}
