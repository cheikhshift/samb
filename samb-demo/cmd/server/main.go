// GENERATED CODE, DO NOT EDIT!
package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/cheikhshift/samb/samb-demo/pkg/api"
)

func main() {

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

func cleanUp(h *http.Server) {
	log.Println("\nShutting down the server...")
	err := h.Close()

	if err != nil {
		panic(err)
	}

	Stop()
	log.Println("Server gracefully stopped")
}
