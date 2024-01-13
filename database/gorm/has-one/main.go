package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
)

type Item struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber // Has one
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type SerialNumber struct {
	ID     int `gorm:"primaryKey"`
	Number string
	ItemID int `gorm:"foreignKey:YourCustomID"` // Has one
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Item{}, &Category{}, &SerialNumber{})

	category := createCategory(db, "Food")

	item := createItem(db, &Item{Name: "Apple", Price: 9.99, CategoryID: category.ID})

	createSerialNumber(db, &SerialNumber{Number: strconv.Itoa(rand.Int()), ItemID: item.ID})

	items := findAllItems(db)
	for _, item := range *items {
		fmt.Println("Item:", item.Name, "with Category:", item.Category.Name, "and Serial Number ID", item.SerialNumber.Number, "found!")
	}

}

func createCategory(db *gorm.DB, name string) *Category {
	var category Category
	// This method fetches the first record matched by the conditions,
	// or creates a new one with the provided value if none was found.
	db.FirstOrCreate(&category, Category{Name: name})
	return &category
}

func createItem(db *gorm.DB, item *Item) *Item {
	db.Create(item)
	return item
}

func createSerialNumber(db *gorm.DB, serialNumber *SerialNumber) *SerialNumber {
	db.Create(serialNumber)
	return serialNumber
}

func findAllItems(db *gorm.DB) *[]Item {
	var items []Item
	db.Preload("Category").Preload("SerialNumber").Find(&items)
	return &items
}
