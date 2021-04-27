package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-service/handlers"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-service", log.LstdFlags)

	// create the handlers
	hh := handlers.NewHello(l)

	// create a new server mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// create a new server
	s := &http.Server{
		Addr:         ":9090",           // Configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write request to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP keep-alive
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	_ = s.Shutdown(tc)
}
