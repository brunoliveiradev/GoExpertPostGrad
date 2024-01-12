package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

type ProductRepository interface {
	InsertProduct(product *Product) error
}

type ProductRepositoryDB struct {
	DB *sql.DB
}

func (p *ProductRepositoryDB) InsertProduct(product *Product) error {
	stmt, err := p.DB.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatalf("Failed to close statement: %v", err)
		}
	}(stmt)

	_, err = stmt.Exec(product.Id, product.Name, product.Price)
	return err
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := createConnectionWithLocalDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 1980.90)

	repo := &ProductRepositoryDB{DB: db}
	err = insertNewProduct(repo, product)
	if err != nil {
		log.Fatalf("Failed to insert product: %v", err)
	}
	fmt.Println("Product inserted successfully with ID: ", product.Id)
}

func createConnectionWithLocalDatabase() (*sql.DB, error) {
	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbName := os.Getenv("DB_NAME")
	//dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	//db, err := sql.Open("mysql", dataSource)
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Database connected and pinged successfully")
	return db, nil
}

func insertNewProduct(repo *ProductRepositoryDB, product *Product) error {
	return repo.InsertProduct(product)
}
