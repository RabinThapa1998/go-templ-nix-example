package main

import (
	"log"
	"net/http"

	"go-templ-nix-example/views"

	"github.com/a-h/templ"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		templ.Handler(views.Index(
			"world",
		)).ServeHTTP(w, r)
	default:
		templ.Handler(views.NotFound()).ServeHTTP(w, r)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerFunc)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
