package main

import "net/http"

func main() {
	// Multiplexer (mux) is a router that receives a request and redirects it to the correct handler
	mux := http.NewServeMux()
	mux.Handle("/hello", site{"Hello, world!"})
	http.ListenAndServe(":8081", mux)
}

type site struct {
	name string
}

func (s site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.name))
}
