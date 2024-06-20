package main 

import "fmt"
import "net/http"
import fs "ascii-art-web/fs"

var not_found = "404 not found"
var not_allowed = "method not allowed"
var max_lenght = 250 


/*making a HandleFunc with a multiplexer*/
var mux = http.NewServeMux()

type Wrapper struct{
	F func(http.ResponseWriter, *http.Request)
}
func (W *Wrapper)ServeHTTP(w http.ResponseWriter, r *http.Request){
	W.F(w,r)
}
func HandleFunc(path string, f func(http.ResponseWriter, *http.Request)){
	wrapper := Wrapper{}	
	wrapper.F = f 
	mux.Handle(path, &wrapper)
}

/* making functions to hanlde requests */
func singlePage(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		w.WriteHeader(404)
		fmt.Fprintln(w , not_found)
		return
	}
	if r.Method == "GET"{
		// return home page  
		fmt.Fprintln(w, "the method used is get")
	}else if r.Method == "POST"{
		// return the template
		fmt.Fprintln(w ,"the method is post")
		test := fs.Ascii_Art("test", "shadow")
		fmt.Fprintln(w, test)
	}else{
		w.WriteHeader(405)
		fmt.Fprintln(w , not_allowed)
	}
}

/* the main function */
func main(){

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux, 
	}
	HandleFunc("/", singlePage)
	fmt.Println("server has been launched at localhost:8080")
	fmt.Println("http://localhost:8080")
	server.ListenAndServe()
}
