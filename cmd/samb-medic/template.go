package main

var recoveryHandlerTemplate = `package api

import (
	"net/http"
)

// %s handles recovery of requests.
func %s(w http.ResponseWriter, r *http.Request, m string) {


	w.Write([]byte("Please implement me."))

}
`
