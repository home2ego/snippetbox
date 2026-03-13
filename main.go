package main

import (
	"log"
	"net/http"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("%sServer starting on http://localhost:4000/%s", green, reset)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
