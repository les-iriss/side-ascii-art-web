package funcs

import (
	"html/template"
	"net/http"

	fs "ascii-art-web/fs"
	"path/filepath"
)


var (
	Index_path        = filepath.Join("templates","index.html")
	Error_path        = filepath.Join("templates","error.html")
	exeeded           = "input exeeded the maximum allowed, try again"
	max_allowed int64 = 50000
	// Data is related to what is shown on the home page
	Data = struct {
		Ascii string
		Err   string
	}{}
	// Input is related to fs function, that is responsible for generating the ascii
	// so these variable are eather sent to that func or recieved from it.
	// the Status and Err are shown on the error page
	Input = struct {
		Text   string
		Banner string
		Status int
		Err    error
	}{
		Status: 200,
		Err:    nil,
	}
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
	wrapper := &Wrapper{} // instence of Wrapper
	wrapper.F = f
	Mux.Handle(path, wrapper)
}

/* making functions to hanlde requests */
func Home(w http.ResponseWriter, r *http.Request) {
	// checking the integrety of the url
	if r.URL.Path != "/" {
		ErrorFunc(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorFunc(w, http.StatusMethodNotAllowed)
	}

	t, err := template.ParseFiles(Index_path)
	if err != nil {
		ErrorFunc(w, http.StatusInternalServerError)
		return
	}
	t.Execute(w, Data)
	Data.Err = ""
	Data.Ascii = ""
}

func Ascii_Art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorFunc(w, http.StatusMethodNotAllowed)
		return
	}
	// checking request length
	Len := r.ContentLength
	if Len > max_allowed {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		t, err := template.ParseFiles(Index_path)
		if err != nil {
			ErrorFunc(w, 500)
			return
		}
		Data.Err = exeeded
		Data.Ascii = ""
		t.Execute(w, Data)
		return
	}

	// parse the form into a map and get the needed value
	r.ParseForm()
	Input.Text = r.FormValue("text")
	Input.Banner = r.FormValue("banner")
	Data.Ascii, Input.Status, Input.Err = fs.Ascii_Art(Input.Text, Input.Banner)
	if Input.Err != nil && Input.Status != 200 {
		ErrorFunc(w, Input.Status)
		return
	}
	http.Redirect(w, r, "/", 303)
}
