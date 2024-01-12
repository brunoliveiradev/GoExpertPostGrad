package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"math/rand"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

type ProductRepository interface {
	InsertProduct(product *Product) error
	UpdateProduct(product *Product) error
	FindProductByID(id string) (*Product, error)
	FindAllProducts() ([]*Product, error)
	DeleteProduct(id string) error
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

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	return err
}

func (p *ProductRepositoryDB) UpdateProduct(product *Product) error {
	stmt, err := p.DB.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepositoryDB) FindProductByID(id string) (*Product, error) {
	stmt, err := p.DB.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product with ID %s not found", id)
		}
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepositoryDB) FindAllProducts() ([]*Product, error) {
	rows, err := p.DB.Query("SELECT id, name, price FROM products LIMIT 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func (p *ProductRepositoryDB) DeleteProduct(id string) error {
	stmt, err := p.DB.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
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

	product := NewProduct("Caneta Azul", rand.Float64())

	repo := &ProductRepositoryDB{DB: db}
	err = insertNewProduct(repo, product)
	if err != nil {
		log.Fatalf("Failed to insert product: %v", err)
	}
	fmt.Printf("Product inserted successfully with name '%s' and price: R$ %.2f\n", product.Name, product.Price)

	product.Price = 1990.90
	err = updateProduct(repo, product)
	if err != nil {
		log.Fatalf("Failed to update product: %v", err)
	}
	fmt.Printf("Product '%s' updated successfully to new price: R$ %.2f\n", product.Name, product.Price)

	product, err = findProductByID(repo, product.ID)
	if err != nil {
		log.Fatalf("Failed to find product: %v", err)
	}
	fmt.Printf("Product '%s' found with price R$ %.2f\n", product.Name, product.Price)

	products, err := findAllProducts(repo)
	if err != nil {
		log.Fatalf("Failed to find products: %v", err)
	}
	fmt.Printf("Found %d products\n", len(products))
	for _, p := range products {
		fmt.Printf("Product '%s' found with price R$ %.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(repo, product.ID)
	if err != nil {
		log.Fatalf("Failed to delete product: %v", err)
	}
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
	fmt.Println("Database connected and pinged successfully!")
	return db, nil
}

func insertNewProduct(repo *ProductRepositoryDB, product *Product) error {
	return repo.InsertProduct(product)
}

func updateProduct(repo *ProductRepositoryDB, product *Product) error {
	return repo.UpdateProduct(product)
}

func findProductByID(repo *ProductRepositoryDB, id string) (*Product, error) {
	return repo.FindProductByID(id)
}

func findAllProducts(repo *ProductRepositoryDB) ([]*Product, error) {
	return repo.FindAllProducts()
}

func deleteProduct(repo *ProductRepositoryDB, id string) error {
	return repo.DeleteProduct(id)
}
