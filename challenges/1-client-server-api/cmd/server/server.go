package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/cmd/server/database"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/pkg/domain"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/util"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	AwesomeApiURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	CotacaoPath   = "/cotacao"
)

func GetCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != CotacaoPath {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	currencyInfo, err := GetLastCurrencyInfoUSDBRL(ctx)
	if err != nil {
		log.Printf("[SERVER][ERROR] Error fetching GetLastCurrencyInfo: %v", err)
		util.HandleError(w, err)
		return
	}

	err = database.SaveCurrencyInfo(ctx, currencyInfo)
	if err != nil {
		log.Printf("[SERVER][ERROR] Error saving CurrencyInfo: %v", err)
		util.HandleError(w, err)
		return
	}

	jsonCurrencyInfo, err := json.MarshalIndent(currencyInfo, "", "  ")
	if err != nil {
		log.Printf("Error formatting currencyInfo as JSON: %v", err)
	} else {
		fmt.Println("[SERVER][INFO] Data saved in database:\n", string(jsonCurrencyInfo))
	}

	w.Header().Set("Content-Type", "application/json")
	serverResponse, err := json.Marshal(currencyInfo.Bid)
	if err != nil {
		util.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(serverResponse)
}

func GetLastCurrencyInfoUSDBRL(ctx context.Context) (*domain.CurrencyInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, AwesomeApiURL, nil)
	if err != nil {
		return nil, err
	}
	log.Println("[SERVER][INFO] Making request to AwesomeApi at", req.URL)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			log.Printf("Request timed out: %v", err)
			return nil, netErr
		}
		return nil, err
	}
	defer resp.Body.Close()

	var currencies *map[string]domain.CurrencyInfo
	if err := json.NewDecoder(resp.Body).Decode(&currencies); err != nil {
		log.Printf("[SERVER][ERROR] Error decoding JSON: %v", err)
		return nil, err
	}
	currency := (*currencies)["USDBRL"]
	log.Println("[SERVER][INFO] AwesomeApi response status:", resp.Status, "response body:", currency)
	return &currency, nil
}
