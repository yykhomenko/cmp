package main

import (
	"fmt"
	"net/http"
)

func main() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminIndex)
	adminMux.HandleFunc("/admin/panic/", panicPage)
	adminHandler := adminAuthMiddleware(adminMux)

	siteMux := http.NewServeMux()
	siteMux.Handle(adminHandler)
	siteMux.HandleFunc("/login", loginPage)
	siteMux.HandleFunc("/logout", logoutPage)
	siteMux.HandleFunc("/", mainPage)

	siteHandler := accessLogHandler(siteMux)
	siteHandler = panicHandler(siteHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", siteHandler)
}
