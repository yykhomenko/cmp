package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yykhomenko/cmp/internal/entity"
	"github.com/yykhomenko/cmp/internal/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService,
	jwtService service.JWTService) LoginController {

	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (lc *loginController) Login(ctx *gin.Context) string {
	var credential entity.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}

	isUserAuthentificated := lc.loginService.
		LoginUser(credential.Email, credential.Password)

	if isUserAuthentificated {
		return lc.jwtService.GenerateToken(credential.Email, true)
	}

	return ""
}
