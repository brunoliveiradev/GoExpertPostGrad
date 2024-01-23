package database

import "github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"

type UserInterface interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
