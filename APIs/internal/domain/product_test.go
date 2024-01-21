package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Smartphone", 1000.00)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, "Smartphone", product.Name)
	assert.Equal(t, 1000.00, product.Price)
}

func TestProduct_NameIsRequiredError(t *testing.T) {
	product, err := NewProduct("", 1000.00)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProduct_PriceIsRequiredError(t *testing.T) {
	product, err := NewProduct("Smartphone", 0)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProduct_PriceInvalidError(t *testing.T) {
	product, err := NewProduct("Smartphone", -1)

	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceInvalid, err)
}

func TestProduct_ValidateProduct(t *testing.T) {
	product, err := NewProduct("Smartphone", 1000.00)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.ValidateProduct())
}
