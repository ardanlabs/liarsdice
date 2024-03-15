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
	http.HandleFunc("/", show)

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

//   go:embed assets/*
//var documents embed.FS

func show(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("").ParseFS(documents, "assets/html/index.html")
	// if err != nil {
	// 	http.Error(w, "Parse: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
	// 	http.Error(w, "Exec:"+err.Error(), http.StatusInternalServerError)
	// }

	d, err := os.ReadFile("app/services/ui/assets/html/index.html")
	if err != nil {
		http.Error(w, "Read: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(d)
}
