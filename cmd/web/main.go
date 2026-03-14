package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCeatePost)

	log.Printf("Server starting on %shttp://localhost:4000%s", "\033[32m", "\033[0m")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
