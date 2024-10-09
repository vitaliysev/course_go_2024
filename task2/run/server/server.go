package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task2/implementation/server"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", server.APIVersionGetRequest)
	mux.HandleFunc("/decode", server.PostRequest)
	mux.HandleFunc("/hard-op", server.HardOpGetRequest)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", ":8080", err)
		}
	}()

	log.Println("Server is ready to handle requests at :8080")
	<-stop
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shut down the server: %v\n", err)
	}
	log.Println("Server stopped")
}
