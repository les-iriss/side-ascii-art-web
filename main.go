package main 

import "fmt"
import "net/http"
import "html/template"
import fs "ascii-art-web/fs"

var not_found = "404 not found"
var not_allowed = "method not allowed"
var internal_error = "internal server error, check your imput"
var exeeded = "the input exeeded the maximul allowed, try again"
var max_allowed int64 = 1000 


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
	len := r.ContentLength
	if len > max_allowed {
				
		t , err := template.ParseFiles("templates/index.html")
		if err != nil {
			t.Execute(w, internal_error)
		}
		t.Execute(w, exeeded)
		
		return
	}
	if r.Method == "GET"{
		t , err := template.ParseFiles("templates/index.html")
		if err != nil {
			// internal server error 
			fmt.Fprintln(w, internal_error)
			return
		}
		t.Execute(w, "")
		// return home page  
	}else if r.Method == "POST"{
		r.ParseForm()
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		Ascii, err := fs.Ascii_Art(text ,banner )
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			t.Execute(w, internal_error)
			return
		}
		t.Execute(w,Ascii)
		// return the template
		if err != nil {
			fmt.Fprintln(w, internal_error)
			fmt.Println(err)
			return
		}
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
