package main

import (
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static/")))
	http.HandleFunc("/about/", handlerIndex)

	http.ListenAndServe(":3000", nil)
}
