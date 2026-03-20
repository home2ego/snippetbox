package main

import (
	"log"
	"net/http"
	"path/filepath"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

// Disable directory listings
func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCeatePost)

	log.Printf("Server starting on %shttp://localhost:4000%s", "\033[32m", "\033[0m")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
