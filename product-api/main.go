package main

import (
	"context"
	"log"
	"my-go-project/product-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProducts(l)

	sm := mux.NewRouter()

	// sm.Handle("/products", productHandler)

	// create a new subrouter for each method with gorilla mux
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.AddProduct)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8080")

		err := server.ListenAndServe()
		if err != nil {
			l.Fatal("Error starting server: ", err)
			os.Exit(1)
		}
	}()

	// graceful shutdown on signal with 30s delay for finishing processes.
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	sig := <-channel
	l.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// deferred function is called when the main function returns
	// we create cancel function to 100% cancel the context and avoid resource leaks.
	defer cancel()
	server.Shutdown(ctx)

}
