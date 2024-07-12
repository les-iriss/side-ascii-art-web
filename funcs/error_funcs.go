package funcs 

import "html/template"
import "net/http"
import "log"

var Error struct{
	Title  string 
	Status  int16
	Hint string
}

var E *template.Template


func init() {
    E, err = template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatal(err,"\n\tAre you Running the Server from a chiled derectory?")
    }
    // Change back to original directory if needed
}
func ErrorFunc(w http.ResponseWriter , Status int){
	switch Status {
	case 404 :
		w.WriteHeader(404)
		Error.Title = "Page Not Found !!!"
		Error.Status = http.StatusNotFound
		Error.Hint = "Page Not Found"
		E.Execute(w , Error)
	case 405 :
		w.WriteHeader(405)
		Error.Title = "Method Not Allowed !!!"
		Error.Status = http.StatusMethodNotAllowed
		Error.Hint = "Method Not Allowed"
		E.Execute(w , Error)
	case 400 : 
		w.WriteHeader(400)
		Error.Title = "Bad Request !!!"
		Error.Status = http.StatusBadRequest
		Error.Hint = "Bad Request"
		E.Execute(w , Error)
	default:
		w.WriteHeader(Status)
		Error.Title = "Internal Server Error !!!"
		Error.Status = http.StatusInternalServerError
		Error.Hint = "InternalServerError"
		E.Execute(w , Error)
	}
}
	
