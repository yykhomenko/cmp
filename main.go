package main

import "net/http"

func main() {
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/admin/", adminIndex)
	adminMux.HandleFunc("/admin/panic/", panicPage)
	adminHandler := adminAuthMiddleware(adminMux)

}
