package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", mainPage)
	r.GET("/login", loginPage)
	r.GET("/logout", logoutPage)
	r.Run()
}

func mainPage(c *gin.Context) {

}

func loginPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

func logoutPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
