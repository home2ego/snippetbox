package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly the home page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func snippetCeatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCeatePost)

	log.Printf("%sServer starting on http://localhost:4000%s", green, reset)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
