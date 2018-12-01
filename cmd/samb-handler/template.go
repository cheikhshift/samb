package main

var verbs = []string{
	"Put",
	"Patch",
	"Delete",
	"Post",
	"Get",
}

var routeTemplate = `package api


var %sNotFound = "{\"error\" : \"Not supported\", \"code\" : 405 }"

// Handle%s Routes PUT,PATCH,DELETE,POST and GET
// requests for corresponding
// handler %s.
func Handle%s(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		%s(w, r)
	case http.MethodPatch:
		%s(w, r)
	case http.MethodDelete:
		%s(w, r)
	case http.MethodPost:
		%s(w, r)
	case http.MethodGet:
		%s(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(%sNotFound))
	}
}

`

var handlerTemplate = `// %s handles %s HTTP requests
// for %s
func %s(w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Please implement me."))

}


`
