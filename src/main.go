package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"BuildingMicroservicesWithGo_NicJackson/src/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// creating handlers
	ph := handlers.NewProducts(l)

	//	creating new server mux and regisering the handlrs
	sm := mux.NewRouter()

	// creating subrouters
	getRouter := sm.Methods("GET").Subrouter()
	putRouter := sm.Methods("PUT").Subrouter()
	postRouter := sm.Methods("POST").Subrouter()

	// registering operations
	getRouter.HandleFunc("/", ph.GetProducts)
	putRouter.HandleFunc("/{id:[0-9]+}", ph.PutProduct)
	postRouter.HandleFunc("/", ph.PostProduct)

	//Creating server object
	s := &http.Server{
		Addr:         "localhost:9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//Starting the server
	go func() {
		l.Println("Starting server on port :9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	//Trapping sigterm or interrupt and gracefully shutting down server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Printf("Received %v signal, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
