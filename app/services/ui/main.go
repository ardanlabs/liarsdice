package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("app/services/ui/website")))

	app := http.Server{
		Addr:    "localhost:8080",
		Handler: http.DefaultServeMux,
	}

	ch := make(chan error, 1)

	go func() {
		ch <- app.ListenAndServe()
	}()

	fmt.Printf("Listening on: %s\n", app.Addr)
	fmt.Println("Hit <ctrl> C to shutdown")
	defer fmt.Println("Shutdown complete")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	select {
	case err := <-ch:
		fmt.Println("ERROR:", err)

	case <-shutdown:
		fmt.Println("\nShutdown requested")
	}

	fmt.Println("Starting shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		app.Close()
	}
}
