package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/login", loginPage)
	r.Run()
}

func loginPage(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("session_id", "yykhomenko")
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
