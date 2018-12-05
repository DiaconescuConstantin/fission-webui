package main

import (
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
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

	http.ListenAndServe(":3000", nil)
}
