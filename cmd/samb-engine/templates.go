package main

var appEngineTemplate string = `// GENERATED CODE, DO NOT EDIT!
package main

import "%s/pkg/api"
import "%s/pkg/hooks"
import "context"
import "google.golang.org/appengine"

func main(){

	hooks.Start()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	http.HandleFunc("/", api.Handler)
	
	go func() {
		<-stop
		cleanUp(h)
	}()

	
	// Launch app on appengine
	appengine.Main()
}

func cleanUp(h *http.Server){
	log.Println("\nShutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	h.Shutdown(ctx)

	hooks.Stop()
	log.Println("Server gracefully stopped")
	
}`

var appEngineConfigTemplate []byte = []byte(`# Generated Code, edit, if not regenerating ;)
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
`)
