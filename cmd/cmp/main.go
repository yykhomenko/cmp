package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yykhomenko/cmp/internal/controller"
	"github.com/yykhomenko/cmp/internal/entity"
	"github.com/yykhomenko/cmp/internal/middleware"
	"github.com/yykhomenko/cmp/internal/repository"
	"github.com/yykhomenko/cmp/internal/service"
)

func main() {
	userRepository := repository.NewUserRepository()
	testUser := entity.User{
		Login:    "yykhomenko",
		Password: "123",
	}
	userRepository.Set(testUser)

	loginService := service.NewLoginService(userRepository)
	jwtService := service.NewJWTService()
	loginController := controller.NewLoginController(loginService, jwtService)

	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})
	authorized := r.Group("/")
	authorized.Use(middleware.AutorizeJWT())
	{
		authorized.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	r.Run()
}
