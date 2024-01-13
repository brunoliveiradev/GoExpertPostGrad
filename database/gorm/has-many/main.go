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
	ID    int `gorm:"primaryKey"`
	Name  string
	Items []Item // Has many relationship
}

type SerialNumber struct {
	ID     int `gorm:"primaryKey"`
	Number string
	ItemID int
}

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	db.AutoMigrate(&Item{}, &Category{}, &SerialNumber{})

	category := createCategory(db, "Car")

	item := createItem(db, &Item{Name: "BMW", Price: 1999.99, Category: *category})

	createSerialNumber(db, &SerialNumber{Number: strconv.Itoa(rand.Int()), ItemID: item.ID})

	items := findAllItems(db)
	for _, item := range *items {
		fmt.Println("Item:", item.Name, "with Category:", item.Category.Name, "and Serial Number", item.SerialNumber.Number, "found!")
	}

	categories := findAllCategories(db)
	for _, category := range *categories {
		fmt.Println("Category", category.Name, "with Items:")
		for _, item := range category.Items {
			fmt.Println("Item:", item.Name, "SN:", item.SerialNumber.Number)
		}
	}
}

func createCategory(db *gorm.DB, name string) *Category {
	var category Category
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

func findAllCategories(db *gorm.DB) *[]Category {
	var categories []Category
	// Preload() is used to load the relationship data
	// Items.SerialNumber is used to load the relationship data of the relationship
	err := db.Model(&Category{}).Preload("Items.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	return &categories
}
