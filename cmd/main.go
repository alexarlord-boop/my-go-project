package main

import (
	"context"
	"log"
	internal "my-go-project/internal/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	// http.ListenAndServe(":8080", sm)

	// creating a custom server with timeouts
	// we can tune params, set timeouts for the server to prevent slow client attacks
	// & is a pointer to the struct -- useful for passing around the struct, not the copy of it
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		// ListenAndServe is a blocking call, so we run it in a goroutine
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// we want to gracefully shutdown the server after receiving a signal from the OS
	// first, we create a channel to receive the signal
	signalChannel := make(chan os.Signal, 1)
	// notify the channel when we receive a signal
	signal.Notify(signalChannel, os.Interrupt, os.Kill)
	// blocks the main goroutine until received a signal
	sig := <-signalChannel
	l.Println("Received terminate signal, graceful shutdown due to: ", sig)

	// creates a new context that will be canceled after 30 seconds. This is useful for setting a deadline for operations to complete.
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// graceful shutdown of the server: wait for the server to finish processing the requests, no new requests.
	s.Shutdown(timeoutContext)

}
