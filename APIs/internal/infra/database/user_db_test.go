package database

import (
	"github.com/brunoliveiradev/courseGoExpert/APIs/internal/domain"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setup(t *testing.T) (*gorm.DB, *User, *domain.User) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&domain.User{})

	userDB := NewUser(db)
	user, _ := domain.NewUser("Saitama", "email@email.com", "password")
	require.NoError(t, userDB.CreateUser(user))

	return db, userDB, user
}

func TestUser_CreateUser(t *testing.T) {
	db, _, user := setup(t)

	var userFound domain.User
	require.NoError(t, db.First(&userFound, "id = ?", user.ID).Error)
	require.Equal(t, user.ID, userFound.ID)
	require.Equal(t, user.Name, userFound.Name)
	require.Equal(t, user.Email, userFound.Email)
	require.NotNil(t, userFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	_, userDB, user := setup(t)

	userFound, err := userDB.FindByEmail(user.Email)
	require.NoError(t, err)
	require.Equal(t, user.ID, userFound.ID)
	require.Equal(t, user.Name, userFound.Name)
	require.Equal(t, user.Email, userFound.Email)
	require.NotNil(t, userFound.Password)
}

func TestUser_FindByEmailNotFoundError(t *testing.T) {
	_, userDB, _ := setup(t)

	_, err := userDB.FindByEmail("email_not_found")
	require.Error(t, err)
	require.Equal(t, ErrUserNotFound, err)
}
