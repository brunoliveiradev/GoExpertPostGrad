package tax

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}
