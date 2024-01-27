package database

import (
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = gorm.ErrRecordNotFound
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *domain.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}
