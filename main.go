package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}
}
