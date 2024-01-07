package main

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}

	reqBody := bytes.NewBuffer([]byte(`{"key":"value"}`))

	resp, err := client.Post("https://httpbin.org/post", "application/json", reqBody)
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
