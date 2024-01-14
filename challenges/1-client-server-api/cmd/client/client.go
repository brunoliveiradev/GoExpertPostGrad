package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func GetCotacao(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			log.Printf("Request timed out: %v", err)
			return netErr
		}
		return err
	}
	defer resp.Body.Close()
	log.Println("Fiz a chamada pro servidor e recebi a resposta:", resp.Status)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return err
	}

	if string(respBody) == "Request timed out" {
		log.Println("O servidor demorou demais pra responder")
		return errors.New("O servidor demorou demais pra responder")
	}

	return saveToFile("challenges/1-client-server-api/output/cotacao.txt", string(respBody))
}

func saveToFile(filename, data string) error {
	// Open the file in append mode, If it does not exist, create it
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file, followed by a newline character
	_, err = fmt.Fprintln(file, "Dólar: "+data)
	return err
}
