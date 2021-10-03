package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/gprisco/nic-series-yt/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api => ", log.LstdFlags)

	ph := handlers.NewProducts(logger)

	sm := mux.NewRouter()

	getSubrouter := sm.Methods(http.MethodGet).Subrouter()
	postSubrouter := sm.Methods(http.MethodPost).Subrouter()
	putSubrouter := sm.Methods(http.MethodPut).Subrouter()

	getSubrouter.HandleFunc("/", ph.GetProducts)

	postSubrouter.HandleFunc("/", ph.AddProduct)

	putSubrouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Put The Server start in a goroutine, so it will not block
	go func() {
		log.Println("Server listening...")
		err := s.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Create a channel
	sigChan := make(chan os.Signal, 1) // 1 -> buffer of size 1

	// Notify sigChan everytime we get Interrupt or Kill Signal
	signal.Notify(sigChan, os.Interrupt)

	// !!! THIS IS BLOCKING -> will listen for signals (specified above by us)
	sig := <-sigChan
	logger.Println("Gracefully shutdown...", sig)

	// Get Background context, assign a 30 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server with the timeout specified above
	s.Shutdown(ctx)
}
