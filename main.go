package main

import (
	"fmt"
	"net/http"

	funcs "my-ascii-art-web/funcs"
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
	//////////////////////////////////
	//serving the static files.
    	fileServer := http.FileServer(http.Dir("./static/"))

	funcs.Mux.Handle("/static/",http.StripPrefix("/static",funcs.MiddleWare(fileServer)))

	//////////////////////////////////
	// handle everything at single page
	funcs.HandleFunc("/", funcs.Home)
	funcs.HandleFunc("/ascii-art", funcs.Ascii_Art)
	funcs.HandleFunc("/download", funcs.Download)
	fmt.Println("server has been launched at localhost:8080")
	fmt.Println("http://localhost:8080")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("\nfatal:\n\tserver has been close. port specified is on use")
		return
	}
}
