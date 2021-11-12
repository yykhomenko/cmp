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

	r.GET("/", mainPage)
	r.GET("/login", loginPage)
	r.Run()
}

func mainPage(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("session_id")
	if id == nil {
		c.String(http.StatusOK, "You need to login %s", `<a href="/login">login</a>`)
	} else {
		c.String(http.StatusOK, "Hello, %s", id)
	}
}

func loginPage(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("session_id", "yykhomenko")
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
