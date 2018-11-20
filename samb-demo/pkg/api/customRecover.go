package api

import "net/http"

func recoverName(w http.ResponseWriter, r *http.Request, n string) {

	println(n + "Logged")
}
