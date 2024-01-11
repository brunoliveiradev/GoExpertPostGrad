package main

import "github.com/google/uuid"

type Product struct {
	Id    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

}
