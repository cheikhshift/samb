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
		cleanUp()
	}()

	
	// Launch app on appengine
	appengine.Main()
}

func cleanUp(){

	hooks.Stop()
	log.Println("App gracefully stopped")
	
}`

var appEngineConfigTemplate []byte = []byte(`# Generated Code, edit, if not regenerating ;)
env: flex
runtime: go
`)
