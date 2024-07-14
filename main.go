package main

import (
	"fmt"
	"net/http"

	funcs "ascii-art-web/funcs"
)

/* the main function */
func main() {
	// initializing the server
	// listening at port 0.0.0.0:8080 
	// so the server can  be accessible from 
	// any network interface on the machine
	// where it is running
	// in contrast to local host where it is accesible only 
	// from 127.0.0.1 ... 
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: funcs.Mux,
	}
	// handle everything at single page
	funcs.HandleFunc("/", funcs.Home)
	funcs.HandleFunc("/ascii-art", funcs.Ascii_Art)
	funcs.HandleFunc("/download", funcs.Download)
	fmt.Println("server has been launched at localhost:8080")
	fmt.Println("http://localhost:8080")
	server.ListenAndServe()
}
