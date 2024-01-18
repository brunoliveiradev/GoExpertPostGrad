package main

import (
	"context"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/cmd/client"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/cmd/server"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/cmd/server/database"
	"log"
	"net/http"
	"time"
)

func main() {
	// Inicializa o banco de dados
	if err := database.InitSqliteDB(); err != nil {
		log.Fatalf("[ERROR] Error initializing database: %v", err)
	}
	defer database.CloseSqliteDB()

	// Inicializa o servidor HTTP em uma goroutine
	go func() {
		http.HandleFunc("/cotacao", server.GetCotacaoHandler)
		log.Println("[DEBUG] Servidor iniciado na porta 8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Aguarda um pouco para garantir que o servidor esteja operacional
	time.Sleep(1 * time.Second)

	// Executa o cliente
	log.Println("[DEBUG] Executando o cliente...")
	if err := client.GetCotacao(context.Background()); err != nil {
		log.Fatalf("[ERROR] Error making request: %v", err)
	}

	log.Println("[DEBUG] Cotação salva com sucesso em 'challenges/1-client-server-api/cmd/output'!")
}
