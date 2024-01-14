package main

import (
	"context"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/cmd/client"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/cmd/server"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/cmd/server/database"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cotacao", server.GetCotacaoHandler)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	if err := database.InitSqliteDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.CloseSqliteDB()

	if err := client.GetCotacao(context.Background()); err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	log.Println("Cotação salva com sucesso!")
}
