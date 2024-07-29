package funcs

import (
	"fmt"
	"html/template"
	fs "my-ascii-art-web/fs"
	"net/http"
	"os"
	"strings"
)

var t *template.Template

var err error

func init() {
	t, err = template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
}

var (
	Index_path        = "templates/index.html"
	Error_path        = "templates/error.html"
	exeeded           = "input exeeded the maximum allowed, try again"
	max_allowed int64 = 1000
	// Data is related to what is shown on the home page
	Data = struct {
		Ascii string
		Err   string
		Shown bool
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
	ToDownload string
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
	t.Execute(w, Data)
	ToDownload = Data.Ascii
	Data.Err = ""
	Data.Ascii = ""
	Data.Shown = false
}

func Ascii_Art(w http.ResponseWriter, r *http.Request) {
	Data.Shown = false
	// this is only for the testing function it is not needed
	// for the the server to operate normally.
	if r.URL.Path != "/ascii-art" {
		ErrorFunc(w, http.StatusNotFound)
	}
	if r.Method != http.MethodPost {
		ErrorFunc(w, http.StatusMethodNotAllowed)
		return
	}
	// checking request length
	Len := r.ContentLength
	if Len > max_allowed {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		Data.Err = exeeded
		Data.Ascii = ""
		t.Execute(w, Data)
		Data.Err = ""
		return
	}

	// parse the form into a map and get the needed value
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		// Handle error
		ErrorFunc(w, 404)
	}
	// Get the needed values from the form
	
	Input.Text = r.PostFormValue("text")
	Input.Banner = r.PostFormValue("banner")
	//Input.Text = r.FormValue("text")
	//Input.Banner = r.FormValue("banner")
	Data.Ascii, Input.Status, Input.Err = fs.Ascii_Art(Input.Text, Input.Banner)
	if Input.Err != nil && Input.Status != 200 {
		ErrorFunc(w, Input.Status)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
	Data.Shown = true
}

func Download(w http.ResponseWriter, r *http.Request) {
	// i have to empty the data after use
	if r.URL.Path != "/download" {
		ErrorFunc(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorFunc(w, http.StatusMethodNotAllowed)
		return
	}
	referer := r.Header.Get("Referer")
	if referer != "http://localhost:8080/" {
		fmt.Println("found error at refere")
		ErrorFunc(w, http.StatusForbidden)
		return
	}
	// creat the file first
	filename := "ascii-art.txt"
	err := os.WriteFile(filename, []byte(ToDownload), 0644)
	if err != nil {
		ErrorFunc(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, filename)
	os.Remove(filename)

}

// using a middleware
func MiddleWare(next http.Handler) http.Handler {
	// the case of examle /something/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			ErrorFunc(w, 404)
			return
		}

		next.ServeHTTP(w, r)
	})
}
