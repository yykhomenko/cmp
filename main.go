package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminIndex)
	adminMux.HandleFunc("/admin/panic/", panicPage)
	adminHandler := adminAuthMiddleware(adminMux)

	siteMux := http.NewServeMux()
	siteMux.Handle("/admin/", adminHandler)
	siteMux.HandleFunc("/login", loginPage)
	siteMux.HandleFunc("/logout", logoutPage)
	siteMux.HandleFunc("/", mainPage)

	siteHandler := accessLogMiddleware(siteMux)
	siteHandler = panicMiddleware(siteHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", siteHandler)
}

func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("panic middleware", r.URL.Path)

		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered", err)
				http.Error(w, "Internal server error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("accessLogMiddleware", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println("[%s] %s, %s %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func adminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("adminAuthMiddleware", r.URL.Path)

		//! auth dummy
		_, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println("no auth at", r.URL.Path)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		//! end of auth dummy

		next.ServeHTTP(w, r)
	})
}

func panicPage(w http.ResponseWriter, r *http.Request) {
	panic("this must be recovered")
}

func adminIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, `<a href="/">site index</a>`)
	fmt.Println(w, "Admin main page")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "yykhomenko",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("sessiod_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusFound)
}
