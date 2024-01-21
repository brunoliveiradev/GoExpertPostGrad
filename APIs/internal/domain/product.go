package domain

import (
	"errors"
	"github.com/brunoliveiradev/courseGoExpert/APIs/pkg/entity"
	"time"
)

var (
	ErrRequiredId       = errors.New("required id")
	ErrProductInvalidID = errors.New("invalid product ID")
	ErrNameIsRequired   = errors.New("name is required")
	ErrPriceIsRequired  = errors.New("price is required")
	ErrPriceInvalid     = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := p.ValidateProduct(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) ValidateProduct() error {
	if p.ID.String() == "" {
		return ErrRequiredId
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrPriceInvalid
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrProductInvalidID
	}
	return nil
}
