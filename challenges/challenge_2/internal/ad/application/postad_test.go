package application

import (
	"errors"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
	"testing"
	"time"
)

func TestPostAdService_Execute(t *testing.T) {
	repo := NewMockRepository(true)
	now := time.Now()
	ad := NewAd(
		NewId("69d39636-d256-47e6-bf86-6bef0cb32ceb"),
		Title{Value: "Test Ad"},
		Description{Value: "This is a test ad"},
		Price{Value: 9.99},
		PublishedAt{Value: now},
	)
	request := PostAdRequest{
		Id:          "69d39636-d256-47e6-bf86-6bef0cb32ceb",
		Title:       "Test Ad",
		Description: "This is a test ad",
		Price:       9.99,
		PublishedAt: now,
	}

	t.Run("success", func(t *testing.T) {
		repo.On("Save", ad).Return(true, nil)
		service := NewPostAdService(repo)

		service.Execute(request)

		repo.AssertCalled(t, "Save", ad)
		assert.NoError(t, nil)
	})

	t.Run("error", func(t *testing.T) {
		expectedError := errors.New("an error")
		repo.On("Save", ad).Return(false, expectedError)
		service := NewPostAdService(repo)

		actualError := service.Execute(request)

		repo.AssertCalled(t, "Save", ad)
		assert.Equal(t, expectedError, actualError)
	})
}
