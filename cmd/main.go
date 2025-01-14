package main

import (
	"log"
	internal "my-go-project/internal/handlers"
	"net/http"
	"os"
)

// creating a simple HTTP server: decomposing the code into smaller parts with handlers and logging
func main() {

	// creating a logger for the data handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	indexHandler := internal.NewIndex(l)
	dataHandler := internal.NewData(l)

	// ServeMux is an HTTP request multiplexer
	sm := http.NewServeMux()
	sm.Handle("/", indexHandler)
	sm.Handle("/data", dataHandler)

	// ListenAndServe uses the DefaultServeMux to handle requests
	// now we are using our own ServeMux with indexHandler and dataHandler registered
	http.ListenAndServe(":8080", sm)
}
