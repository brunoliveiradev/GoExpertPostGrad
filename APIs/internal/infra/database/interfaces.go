package database

import "github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"

type UserInterface interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

type ProductInterface interface {
	Create(product *domain.Product) error
	FindAll(page int, limit int, sort string) ([]*domain.Product, error)
	FindByID(id string) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id string) error
}
