package funcs

import (
	fs "ascii-art-web/fs"
	"fmt"
	"html/template"
	"net/http"
)

var (
	template_path        = "templates/index.html"
	not_found            = "404 not found"
	not_allowed          = "405 Method Not Allowed"
	internal_error       = "500 Internal Server Error, error check your imput"
	exeeded              = "413 input exeeded the maximum allowed, try again"
	max_allowed    int64 = 1000
)

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
func Home(w http.ResponseWriter, r *http.Request) {
	// checking the integrety of the url
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte(not_found))
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte(not_allowed))
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// internal server error
		fmt.Fprintln(w, internal_error)
		return
	}
	t.Execute(w, "")
}

func Ascii_Art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte(not_allowed))
		return
	}
	// checking request length
	Len := r.ContentLength
	if Len > max_allowed {
		w.WriteHeader(413)
		t, err := template.ParseFiles(template_path)
		if err != nil {
			t.Execute(w, internal_error)
		}
		t.Execute(w, exeeded)
		return
	}
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
}
