package database

import (
	"fmt"
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func setupProduct(t *testing.T) (*gorm.DB, *Product, *domain.Product) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.Product{})

	productDB := NewProduct(db)
	product, _ := domain.NewProduct("Smartphone", 1299.99)
	require.NoError(t, productDB.Create(product))

	return db, productDB, product
}

func TestProduct_Create(t *testing.T) {
	db, _, product := setupProduct(t)

	var productFound domain.Product
	require.NoError(t, db.First(&productFound, "id = ?", product.ID).Error)
	assert.NotEqual(t, "", productFound.ID)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProduct_FindAll(t *testing.T) {
	db, productDB, _ := setupProduct(t)

	for i := 0; i < 10; i++ {
		product, _ := domain.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		require.NoError(t, productDB.Create(product))
	}

	products, err := productDB.FindAll(1, 10, "asc")
	require.NoError(t, err)
	assert.Len(t, products, 10)

	var productFound domain.Product
	require.NoError(t, db.First(&productFound, "id = ?", products[0].ID).Error)
	assert.Equal(t, products[0].ID, productFound.ID)
	assert.Equal(t, products[0].Name, productFound.Name)
	assert.Equal(t, products[0].Price, productFound.Price)

	products, err = productDB.FindAll(2, 10, "asc")
	require.NoError(t, err)
	assert.Len(t, products, 1)

	products, err = productDB.FindAll(3, 10, "asc")
	require.NoError(t, err)
	assert.Len(t, products, 0)
}

func TestProduct_FindByID(t *testing.T) {
	db, productDB, product := setupProduct(t)

	productFound, err := productDB.FindByID(product.ID.String())
	require.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)

	var productNotFound domain.Product
	err = db.First(&productNotFound, "id = ?", "not_found").Error
	assert.Error(t, err)
	assert.Equal(t, ErrProductNotFound, err)
}

func TestProduct_Update(t *testing.T) {
	db, productDB, product := setupProduct(t)

	product.Name = "New Name"
	product.Price = 999.99
	require.NoError(t, productDB.Update(product))

	var productFound domain.Product
	require.NoError(t, db.First(&productFound, "id = ?", product.ID).Error)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProduct_UpdateProductNotFound(t *testing.T) {
	_, productDB, _ := setupProduct(t)

	product, _ := domain.NewProduct("New Something", 1299.99)
	product.Name = "New Name"
	product.Price = 999.99

	err := productDB.Update(product)
	assert.Error(t, err)
	assert.Equal(t, ErrProductNotFound, err)
}

func TestProduct_Delete(t *testing.T) {
	db, productDB, product := setupProduct(t)

	require.NoError(t, productDB.Delete(product.ID.String()))

	var productFound domain.Product
	err := db.First(&productFound, "id = ?", product.ID).Error
	assert.Error(t, err)
	assert.Equal(t, ErrProductNotFound, err)
}
