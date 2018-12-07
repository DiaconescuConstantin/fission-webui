package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	log.Println("Starting server...")

	changeHeaderThenServe := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Set some header.
			w.Header().Add("Keep-Alive", "300")
			// Serve with the actual handler.
			h.ServeHTTP(w, r)
		}
	}
	http.Handle("/", changeHeaderThenServe(http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/about/", handlerIndex)
	http.HandleFunc("/api-node/", handlerIndex)

	log.Println("Listening on port 3000...")

	s := http.Server{Addr: ":3000"}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
