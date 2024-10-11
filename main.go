package main

import (
	"context"
	"domolitom/microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	hh := handlers.NewHello(log.New(os.Stdout, "product-api", log.LstdFlags))
	gh := handlers.NewBye(log.New(os.Stdout, "product-api", log.LstdFlags))
	th := handlers.NewTestHandler(log.New(os.Stdout, "product-api", log.LstdFlags))
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", gh)
	sm.Handle("/test/test/test", th)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	log.Println("Starting Server")

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(tCtx)

}
