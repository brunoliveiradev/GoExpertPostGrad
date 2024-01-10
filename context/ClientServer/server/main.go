package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request began")
	defer log.Println("request ended")

	select {
	case <-time.After(5 * time.Second):
		log.Println("waited for 5 seconds")
		w.Write([]byte("Request processed"))
	case <-ctx.Done():
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}

}
