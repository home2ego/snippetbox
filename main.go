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
	w.Write([]byte("Dispaly the home page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("%sServer starting on http://localhost:4000/%s", green, reset)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
