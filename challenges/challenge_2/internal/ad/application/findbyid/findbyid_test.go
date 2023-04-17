package findbyid

import (
	"errors"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
	"testing"
	"time"
)

func Test_findAdByIdService_Execute(t *testing.T) {
	now := time.Now()
	expected, _ := NewAd(
		"123",
		"Test Ad",
		"This is a test expected",
		9.99,
		now,
	)
	t.Run("success", func(t *testing.T) {
		repo := NewMockRepository(true)
		repo.On("FindBy", expected.Id()).Return(&expected, nil)
		service := NewFindAdByIdService(repo)

		actual, _ := service.Execute(NewFindAdByIdRequest(expected.Id().String()))

		assert.Equal(t, &expected, actual)
	})

	t.Run("not found", func(t *testing.T) {
		repo := NewMockRepository(false)
		repo.On("FindBy", expected.Id()).Return(nil, nil)
		service := NewFindAdByIdService(repo)

		actual, _ := service.Execute(NewFindAdByIdRequest(expected.Id().String()))

		assert.Nil(t, actual)
	})

	t.Run("error", func(t *testing.T) {
		expectedError := errors.New("an error")
		repo := NewMockRepository(false)
		repo.On("FindBy", expected.Id()).Return(nil, expectedError)
		service := NewFindAdByIdService(repo)

		_, actualError := service.Execute(NewFindAdByIdRequest(expected.Id().String()))

		assert.Equal(t, expectedError, actualError)
	})
}
