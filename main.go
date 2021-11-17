package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yykhomenko/cmp/controller"
	"github.com/yykhomenko/cmp/middleware"
	"github.com/yykhomenko/cmp/repository"
	"github.com/yykhomenko/cmp/service"
)

func main() {
	userRepository := repository.NewUserRepository()
	loginService := service.NewLoginService(userRepository)
	jwtService := service.NewJWTService()
	loginController := controller.NewLoginController(loginService, jwtService)

	r := gin.Default()
	r.Use(middleware.AutorizeJWT())
	r.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})

	r.Run()
}
