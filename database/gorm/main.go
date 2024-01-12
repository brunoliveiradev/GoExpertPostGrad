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
