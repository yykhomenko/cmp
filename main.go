package main

import (
	"fmt"
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
	r.GET("/logout", logoutPage)
	r.Run()
}

func mainPage(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("session_id")
	if id == nil {
		c.Data(http.StatusOK, "text/html; charset=utf-8",
			[]byte(`You need to login <a href="/login">login</a>`))
	} else {
		c.Data(http.StatusOK, "text/html; charset=utf-8",
			[]byte(fmt.Sprintf("Hello, %s %s", id, `<a href="/logout">logout</a>`)))
	}
}

func loginPage(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("session_id", "yykhomenko")
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func logoutPage(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
