package domain

import (
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
	filledWithData bool
}

func NewMockRepository(filledWithData bool) *mockRepository {
	return &mockRepository{filledWithData: filledWithData}
}

func (m *mockRepository) FindBy(id Id) (*Ad, error) {
	args := m.Called(id)

	if m.filledWithData {
		return args.Get(0).(*Ad), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (m *mockRepository) FindSetOf(number int) ([]Ad, error) {
	args := m.Called(number)

	if m.filledWithData {
		return args.Get(0).([]Ad), args.Error(1)
	} else {
		return []Ad{}, args.Error(1)
	}
}

func (m *mockRepository) Save(ad Ad) (bool, error) {
	args := m.Called(ad)
	return args.Bool(0), args.Error(1)
}
