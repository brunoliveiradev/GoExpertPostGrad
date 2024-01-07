package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	FirstExample()
	SecondExample()
	ThirdExample()
	FourthExample()
}

// FirstExample function demonstrates a simple HTTP POST request using the http.Client.
func FirstExample() {
	client := &http.Client{Timeout: 5 * time.Second}

	// Create a new request body
	reqBody := bytes.NewBuffer([]byte(`{"key":"value"}`))

	// Send a POST request to the specified URL with the created request body
	resp, err := client.Post("https://httpbin.org/post", "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read all data from the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(respBody))
}

// SecondExample function demonstrates a simple HTTP GET request using the http.Client and Do.
func SecondExample() {
	// Create a new HTTP client with a request timeout of 5 seconds
	c := &http.Client{Timeout: 5 * time.Second}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	// If an error occurred while creating the request, panic
	if err != nil {
		panic(err)
	}
	// Set the Accept header of the request to "application/json"
	req.Header.Set("Accept", "application/json")

	// Send the HTTP request
	resp, err := c.Do(req)
	// If an error occurred during the request, panic
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(respBody))
}

// ThirdExample function demonstrates a simple HTTP GET request using the http.Client and context.
func ThirdExample() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

// FourthExample function demonstrates a simple HTTP GET request using the http.Client and context returning a Timeout.
func FourthExample() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
