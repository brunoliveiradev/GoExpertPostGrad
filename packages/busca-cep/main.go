package main

import (
	"encoding/json"
	"errors"
	"github.com/brunoliveiradev/GoExpertPostGrad/packages/busca-cep/domain"
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"time"
)

const (
	getPath   = "/busca"
	getKey    = "cep"
	localPort = ":8080"
)

var httpClient = &http.Client{Timeout: 5 * time.Second}

func main() {
	http.HandleFunc(getPath, BuscaCepHandler)
	http.ListenAndServe(localPort, nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != getPath {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get(getKey)
	if cepParam == "" {
		http.Error(w, "CEP is required", http.StatusBadRequest)
		return
	}

	matched, _ := regexp.MatchString(`^\d{8}$`, cepParam)
	if !matched {
		http.Error(w, "Invalid CEP format", http.StatusBadRequest)
		return
	}

	cep, err := BuscaCep(cepParam)
	if err != nil {
		log.Printf("Error fetching CEP: %v", err)

		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			http.Error(w, "Request timed out", http.StatusRequestTimeout)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := json.Marshal(cep)
	if err != nil {
		log.Printf("Error encoding CEP: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func BuscaCep(cep string) (*domain.ViaCEP, error) {
	resp, err := httpClient.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			log.Printf("Request timed out: %v", err)
			return nil, netErr
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var viaCEP domain.ViaCEP
	err = json.Unmarshal(body, &viaCEP)
	if err != nil {
		return nil, err
	}

	return &viaCEP, nil
}
