package transpiler


// cmdWrapper is used to generate go source
// within the generated program package.
var cmdWrapper string = `// GENERATED CODE, DO NOT EDIT!
package main

import (
		%s
)

func %s(){
	%s
}`

// globalWrapper is used to generate files
// for package global.
var globalWrapper string = `// Package globals has your applications
// global variables, exported as package identifiers
// GENERATED CODE, DO NOT EDIT!
package globals

%s`


// configWrapper is used to generate the file
// with a web server's port and host information
// as a variables.
var configWrapper string = `// GENERATED CODE, DO NOT EDIT!
package main

var port = "%s"
var host = "%s"
var webroot = "%s"`

// routeWrapper used to generate route handler function.
var routeWrapper string = `//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		%s
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			defer catchPanic(w,r)

			%s
}`

// recoveryWrapper is used to generate source file
// for package api's recover function.
var recoveryWrapper string = `package api

// Function used to get
// error message related 
// to panic.
func catchPanic(w http.ResponseWriter, r * http.Request){

	if n := recover(); n != nil {

		%s

	}
}`

// mainWrapper is used to generate the source for your 
// web server. This launches the generated server source.
var mainWrapper string = `// GENERATED CODE, DO NOT EDIT!
package main

import "%s/pkg/api"
import "context"

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
		log.Println(err)
	}

	
}

func cleanUp(h *http.Server){
	log.Println("\nShutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	h.Shutdown(ctx)

	Stop()
	log.Println("Server gracefully stopped")
	
}`
