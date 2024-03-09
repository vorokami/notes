// server starter
package main

import (
	"fmt"
	"log"
	"net/http"
)

var spec = ":8090" // localhost:8090

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(http.CanonicalHeaderKey("content-type"), "text/plain")
		w.Write([]byte("got path"))
	})

	mux.HandleFunc("/note/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "showing note with id=%v\n", id)
	})

	if err := http.ListenAndServe(spec, mux); err != nil {
		log.Fatalf("Failed to start server on %s: %v", spec, err)
	}
}