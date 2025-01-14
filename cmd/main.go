package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// creating a simple HTTP server
func main() {
	// HandleFunc takes function, creates HTTP handler and adds it to the DefaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
	})

	http.HandleFunc("/another", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Another Request received")
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		data, error := io.ReadAll(r.Body)
		if error != nil {
			http.Error(w, "Error reading data", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "User, Your input was: %s\n", string(data))
		log.Println("Data received: ", string(data))
	})

	// ListenAndServe uses the DefaultServeMux to handle requests
	http.ListenAndServe(":8080", nil)
}
