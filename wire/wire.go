//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/brunoliveiradev/GoExpertPostGrad/wire/product"
	"github.com/google/wire"
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
