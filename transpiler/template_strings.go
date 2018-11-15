package transpiler

var cmdWrapper string = `// GENERATED CODE, DO NOT EDIT!
package main

func %s(){
	%s
}`

var globalWrapper string = `// Package globals has your applications
// global variables, exported as package identifiers
// GENERATED CODE, DO NOT EDIT!
package globals

%s`

var configWrapper string = `// GENERATED CODE, DO NOT EDIT!
package main

var port = "%s"
var host = "%s"
var webroot = "%s"`

var routeWrapper string = `//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		%s
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			%s
}`


var mainWrapper string = `// GENERATED CODE, DO NOT EDIT!
package main

import "%s/pkg/api"

func main(){

	Start()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	http.HandleFunc("/", api.Handler)

	h := &http.Server{Addr: host + ":" + port}

	go func() {
		<-stop
		cleanUp(h)
	}()

	err := h.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func cleanUp(h *http.Server){
	log.Println("\nShutting down the server...")
	err := h.Close()

	if err != nil {
		panic(err)
	}

	Stop()
	log.Println("Server gracefully stopped")
}`
