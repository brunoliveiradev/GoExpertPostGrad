package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Items []Item // Has many relationship
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Item{}, &Category{})

	category := createCategory(db, "Car")

	createItem(db, &Item{Name: "BMW", Price: 999999.99, CategoryID: category.ID})

	items := findAllItems(db)
	for _, item := range *items {
		fmt.Println("Item:", item.Name, "with Category:", item.Category.Name, "found!")
	}

	categories := findAllCategories(db)
	for _, category := range *categories {
		fmt.Println("Category", category.Name, "with Items:")
		for _, item := range category.Items {
			fmt.Println("Item:", item.Name)
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

func findAllItems(db *gorm.DB) *[]Item {
	var items []Item
	db.Preload("Category").Find(&items)
	return &items
}

func findAllCategories(db *gorm.DB) *[]Category {
	var categories []Category
	err := db.Model(&Category{}).Preload("Items").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	return &categories
}
