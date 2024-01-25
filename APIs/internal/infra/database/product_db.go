package database

import (
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrProductNotFound = gorm.ErrRecordNotFound
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *domain.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindByID(id string) (*domain.Product, error) {
	var product domain.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, ErrProductNotFound
	}
	return &product, nil
}

func (p *Product) Update(product *domain.Product) error {
	// Save updates value in database. If value doesn't contain a matching primary key, value is inserted as a new record.
	// that's why we need to check if the product exists
	if _, err := p.FindByID(product.ID.String()); err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}

func (p *Product) FindAll(page int, limit int, sort string) ([]*domain.Product, error) {
	var products []*domain.Product

	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page == 0 && limit == 0 {
		page = 1
		limit = 10
	}

	err := p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error

	return products, err
}
