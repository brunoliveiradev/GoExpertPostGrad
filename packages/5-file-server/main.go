package main

import (
	"log"
	"net/http"
)

func main() {
	// serve static files
	fileServer := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello again!"))
	})

	log.Fatalf("Error starting server: %v", http.ListenAndServe(":8080", mux))
}
