package main

import (
	"io"
	"net/http"
)

func main() {

	// GET request
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	// Read the response body into a byte array and keep it in memory
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
	req.Body.Close()
}
