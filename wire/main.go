package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal("error opening database: ", err)
	}

	useCase := NewProductUseCase(db)

	productOutput, err := useCase.GetProductByID(1)
	if err != nil {
		log.Fatal("error getting product: ", err)
	}

	log.Println("Product:", productOutput.ID, productOutput.Name)
}
