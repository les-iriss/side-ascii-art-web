package funcs

import fs "ascii-art-web/fs"
import "fmt"
import "html/template"
import "net/http"

var template_path = "templates/index.html"
var not_found = "404 not found"
var not_allowed = "405 Method Not Allowed"
var internal_error = "500 Internal Server Error, error check your imput"
var exeeded = "413 input exeeded the maximum allowed, try again"
var max_allowed int64 = 1000

/*making a HandleFunc with a multiplexer*/
var Mux = http.NewServeMux()

type Wrapper struct {
	F func(http.ResponseWriter, *http.Request)
}

// giving the Wrapper type a ServerHTTP method to satisfy
// the Handler interface
func (W *Wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	W.F(w, r)
}

// making a HandleFunc to wrap functions into handlers
func HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	wrapper := Wrapper{}
	wrapper.F = f
	Mux.Handle(path, &wrapper)
}

/* making functions to hanlde requests */
func SinglePage(w http.ResponseWriter, r *http.Request) {
	// checking the integrety of the url
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		fmt.Fprintln(w, not_found)
		return
	}
	// checking request length
	len := r.ContentLength
	if len > max_allowed {
		w.WriteHeader(413)
		t, err := template.ParseFiles(template_path)
		if err != nil {
			t.Execute(w, internal_error)
		}
		t.Execute(w, exeeded)
		return
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			// internal server error
			fmt.Fprintln(w, internal_error)
			return
		}
		t.Execute(w, "")
	} else if r.Method == "POST" {
		// parse the form into a map and get the needed value
		r.ParseForm()
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		Ascii, err1 := fs.Ascii_Art(text, banner)
		t, err2 := template.ParseFiles("templates/index.html")
		if err1 != nil || err2 != nil {
			t.Execute(w, internal_error)
			return
		}
		t.Execute(w, Ascii)
	} else {
		w.WriteHeader(405)
		fmt.Fprintln(w, not_allowed)
	}
}
