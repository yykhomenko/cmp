package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yykhomenko/cmp/entity"
	"github.com/yykhomenko/cmp/service"
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

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential entity.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}

	isUserAuthentificated := controller.loginService.
		LoginUser(credential.Email, credential.Password)

	if isUserAuthentificated {
		return controller.jwtService.
			GenerateToken(credential.Email, true)
	}

	return ""
}