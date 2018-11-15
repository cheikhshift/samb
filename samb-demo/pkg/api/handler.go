//package api contains your web app's handler definitions.
// GENERATED CODE, DO NOT EDIT!!
package api

import (
	"fmt"
	"net/http"
	"strings"
)

// Handles routing of application.
func Handler(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.Path, "/hello/") {
		println("Hello")

		if strings.Contains(r.URL.Path, "/hello/Foo") && r.Method == "GET" {
			println("Hello")
			fmt.Println("Hello")
		}
	}
}
