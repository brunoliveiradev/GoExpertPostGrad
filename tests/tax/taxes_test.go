package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTaxes(t *testing.T) {
	assert.Equal(t, 0.0, CalculateTaxes(-1))
	assert.Equal(t, 0.0, CalculateTaxes(0))
	assert.Equal(t, 25.0, CalculateTaxes(500))
	assert.Equal(t, 100.0, CalculateTaxes(1000))
	assert.Equal(t, 2500.0, CalculateTaxes(10000))
}

func TestCalculateTaxesAndSave(t *testing.T) {
	repo := &RepositoryMock{}
	repo.On("SaveTax", 15.0).Return(nil).Once()
	repo.On("SaveTax", 0.0).Return(errors.New("negative amount"))

	err := CalculateTaxesAndSave(repo, 1001.0)
	assert.Nil(t, err)

	err = CalculateTaxesAndSave(repo, 0.0)
	assert.Error(t, err, "negative amount")

	repo.AssertExpectations(t)
	repo.AssertNumberOfCalls(t, "SaveTax", 2)
}
