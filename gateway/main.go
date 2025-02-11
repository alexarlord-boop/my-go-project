package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"os/signal"
	"syscall"
	"time"

	protos "my-go-project/protos/currency"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 1. create a grpc currency client by using code generated from protos/currency.proto
	// create a new client connection to the grpc server (same host, port)
	cc, err := grpc.NewClient("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Print("Unable to connect to currency service", "error", err)
		os.Exit(1)
	}
	defer cc.Close() // close the connection when the main function exits

	// create a new currency service client
	currencyServiceClient := protos.NewCurrencyServiceClient(cc) // create a new currency service client

	// 2. define the gateway (http server)
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	sm := mux.NewRouter()
	// create a new subrouter for http method with gorilla mux
	getR := sm.Methods(http.MethodGet).Subrouter()
	// create a new route for /products that will handle a request with invoking grpc request
	getR.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		rateRequest := &protos.RateRequest{
			// Base:        protos.Currencies_USD,    .proto enum generates diffrent types for mapping
			// Destination: protos.Currencies_EUR,
			Base:        protos.Currencies(protos.Currencies_value["USD"]),
			Destination: protos.Currencies(protos.Currencies_value["EUR"]),
		}

		currencyRate, err := currencyServiceClient.GetRate(context.Background(), rateRequest)

		if err != nil {
			log.Print("Unable to get rate", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("Rate: %v", currencyRate)
		w.Write([]byte(
			"Hello from the gateway! The rate is " + fmt.Sprintf("%.2f", currencyRate.Rate),
		))
	})

	// CORS handler for allowing requests from localhost:3000
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler(sm),
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 3. start the server
	go func() {
		l.Println("Starting api gateway server on port 8080")

		err := server.ListenAndServe()
		if err != nil {
			l.Fatal("Error starting server: ", err)
			os.Exit(1)
		}
	}()

	// 4. graceful shutdown
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	sig := <-channel
	l.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
