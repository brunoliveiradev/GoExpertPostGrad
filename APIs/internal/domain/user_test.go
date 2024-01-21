package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Saitama", "saitama@email.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Saitama", user.Name)
}

func TestUser_ValidatedPassword(t *testing.T) {
	user, err := NewUser("Saitama", "saitama@email.com", "13131313")

	assert.Nil(t, err)
	assert.True(t, user.ValidatedPassword("13131313"))
	assert.False(t, user.ValidatedPassword("13131314"))
	assert.NotEqual(t, "13131313", user.Password)
}
