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
	Categories []Category `gorm:"many2many:itens_categories;"` // many2many relationship
	gorm.Model
}

type Category struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Items []Item `gorm:"many2many:itens_categories;"` // many2many relationship
}

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	db.AutoMigrate(&Item{}, &Category{})

	category := createCategory(db, "Clothes")
	otherCategory := createCategory(db, "Sports")

	createItem(db, &Item{Name: "T Shirt", Price: 19.99, Categories: []Category{*category, *otherCategory}})

	itemsAndCategories := findAllItemsAndCategories(db)
	for _, category := range *itemsAndCategories {
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

func findAllItemsAndCategories(db *gorm.DB) *[]Category {
	var categories []Category
	err := db.Model(&Category{}).Preload("Items").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	return &categories
}
