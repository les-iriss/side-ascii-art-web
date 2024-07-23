package funcs 

import "html/template"
import "net/http"

var Error struct{
	Title  string 
	Status  int16
	Hint string
}

var E *template.Template


func init() {
    E, err = template.ParseFiles("templates/error.html")
    if err != nil {
        panic(err)
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
	case http.StatusForbidden:
		w.WriteHeader(http.StatusForbidden)
		Error.Title = "Forbidden !!!"
		Error.Status = http.StatusForbidden
		Error.Hint = "the way you access page download is not supported"
	default:
		w.WriteHeader(Status)
		Error.Title = "Internal Server Error !!!"
		Error.Status = http.StatusInternalServerError
		Error.Hint = "InternalServerError"
		E.Execute(w , Error)
	}
}
	
