package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Item struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Item{}, &Category{})

	category := createCategory(db, "Electronics")

	item := &Item{Name: "Notebook", Price: 20000, CategoryID: category.ID}

	item = createItem(db, item)
	fmt.Println("Item:", item.Name, "with Category:", category.Name, "created successfully!")

	items := findAllItems(db)
	for _, item := range *items {
		fmt.Println("Item:", item.Name, "with Category:", item.Category.Name)
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

func findAllItems(db *gorm.DB) *[]Item {
	var items []Item
	db.Preload("Category").Find(&items)
	return &items
}
