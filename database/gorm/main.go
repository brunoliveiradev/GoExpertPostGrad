package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uuid.UUID `gorm:"primary_key"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//createProducts(db)
	//findProducts(db)
	findWithLimitAndOffset(db, 3, 1)
	findWithWereByPrice(db, 2000)
	findWithLike(db, "%Azul Caneta%")
	updateProductPriceGivenID(db, uuid.MustParse("cef49218-abd9-41af-b26a-6a0320f9c8d3"), 1111)
	//deleteProductGivenID(db, uuid.MustParse("b46094f0-0d27-44d6-b101-1d270d9d2f24"))
}

func createProducts(db *gorm.DB) {
	db.AutoMigrate(&Product{})

	products := []Product{
		{
			ID:    uuid.New(),
			Name:  "Notebook",
			Price: 2000,
		},
		{
			ID:    uuid.New(),
			Name:  "Tablet",
			Price: 3000,
		},
		{
			ID:    uuid.New(),
			Name:  "Smartphone",
			Price: 4000,
		},
	}
	db.Create(&products)
}

func findProducts(db *gorm.DB) {
	// select one
	var product Product
	db.First(&product, "name = ?", "Caneta Azul")
	fmt.Println(product)

	// select all
	var products []Product
	db.Find(&products)

	for _, p := range products {
		fmt.Println(p)
	}
}

func findWithLimitAndOffset(db *gorm.DB, limit, offset int) {
	fmt.Println("Products with limit", limit, "and offset", offset)
	var products []Product
	db.Limit(limit).Offset(offset).Find(&products)

	for _, p := range products {
		fmt.Println(p)
	}
}

func findWithWereByPrice(db *gorm.DB, price float64) {
	fmt.Println("Products with price greater than", price)
	var products []Product
	db.Where("price > ?", price).Find(&products)

	for _, p := range products {
		fmt.Println(p)
	}
}

func findWithLike(db *gorm.DB, name string) {
	fmt.Println("Products with name like", name)
	var products []Product
	db.Where("name LIKE ?", name).Find(&products)

	for _, p := range products {
		fmt.Println(p)
	}
}

func updateProductPriceGivenID(db *gorm.DB, id uuid.UUID, newPrice float64) {
	var p Product
	db.First(&p, "id = ?", id)

	p.Price = newPrice
	db.Save(&p)

	fmt.Println("Product updated", p)
}

func deleteProductGivenID(db *gorm.DB, id uuid.UUID) {
	var p Product
	db.First(&p, "id = ?", id)
	fmt.Println("Product deleted", p)
	db.Delete(&p)
}
