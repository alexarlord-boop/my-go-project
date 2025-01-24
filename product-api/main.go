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

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProducts(l)

	sm := mux.NewRouter()

	// sm.Handle("/products", productHandler)

	// create a new subrouter for each method with gorilla mux
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", productHandler.GetList)
	getR.HandleFunc("/products/{id:[0-9]+}", productHandler.GetDetails)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/products/{id:[0-9]+}", productHandler.Update)
	putR.Use(productHandler.MiddlewareValidateProduct)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/products", productHandler.Create)
	postR.Use(productHandler.MiddlewareValidateProduct)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

	// redoc options (redoc - swagger ui alternative)
	swaggerDocOptions := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerHandler := middleware.Redoc(swaggerDocOptions, nil)

	// ui with redoc. we also need endpoint to serve the swagger.yaml file
	getR.Handle("/docs", swaggerHandler)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS handler for allowing requests from localhost:3000
	// "*" is wildcard for all origins, for example, we can use it for testing purposes, or for public APIs.
	// if you require the client to pass authentication headers (e.g. cookies) you need to specify the origin -- the value can not be *.
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler(sm),
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
