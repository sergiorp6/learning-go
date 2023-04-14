package application

import (
	"errors"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
	"testing"
	"time"
)

func Test_findAdByIdService_Execute(t *testing.T) {
	now := time.Now()
	expected := NewAd(
		NewId("123"),
		Title{Value: "Test Ad"},
		Description{Value: "This is a test expected"},
		Price{Value: 9.99},
		PublishedAt{Value: now},
	)
	t.Run("success", func(t *testing.T) {
		repo := NewMockRepository(true)
		repo.On("FindBy", expected.Id()).Return(&expected, nil)
		service := NewFindAdByIdService(repo)

		actual, _ := service.Execute(FindAdByIdRequest{Id: expected.Id().String()})

		assert.Equal(t, &expected, actual)
	})

	t.Run("not found", func(t *testing.T) {
		repo := NewMockRepository(false)
		repo.On("FindBy", expected.Id()).Return(nil, nil)
		service := NewFindAdByIdService(repo)

		actual, _ := service.Execute(FindAdByIdRequest{Id: expected.Id().String()})

		assert.Nil(t, actual)
	})

	t.Run("error", func(t *testing.T) {
		expectedError := errors.New("an error")
		repo := NewMockRepository(false)
		repo.On("FindBy", expected.Id()).Return(nil, expectedError)
		service := NewFindAdByIdService(repo)

		_, actualError := service.Execute(FindAdByIdRequest{Id: expected.Id().String()})

		assert.Equal(t, expectedError, actualError)
	})
}
