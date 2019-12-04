package main

import (
	"context"
	"feed-time-calculation/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// устанвливаем обратку системного сигнала
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	feedTimeApi := api.InitApi()

	router := http.NewServeMux()
	router.HandleFunc("/feed-time", feedTimeApi.GetFeedTime)

	addr := ":8080"

	log.Println("Starting server on", addr)

	server := &http.Server{Addr: addr, Handler: router}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", addr, err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shotdown error: %v\n", err)
	}

	log.Println("Server stopped")

}
