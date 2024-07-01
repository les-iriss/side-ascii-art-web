package main

import (
	"fmt"
	"net/http"

	funcs "ascii-art-web/funcs"
)

/* the main function */
func main() {
	// initializing the server
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: funcs.Mux,
	}
	// handle everything at single page
	funcs.HandleFunc("/", funcs.Home)
	funcs.HandleFunc("/ascii-art", funcs.Ascii_Art)
	fmt.Println("server has been launched at localhost:8080")
	fmt.Println("http://localhost:8080")
	server.ListenAndServe()
}
