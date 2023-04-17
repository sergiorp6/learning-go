package postad

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
	"testing"
	"time"
)

func TestPostAdService_Execute(t *testing.T) {
	now := time.Now()
	ad, _ := NewAd(
		"69d39636-d256-47e6-bf86-6bef0cb32ceb",
		"Test Ad",
		"This is a test ad",
		9.99,
		now,
	)
	request := PostAdRequest{
		"69d39636-d256-47e6-bf86-6bef0cb32ceb",
		"Test Ad",
		"This is a test ad",
		9.99,
		now,
	}

	t.Run("success", func(t *testing.T) {
		repo := NewMockRepository(true)
		repo.On("Save", ad).Return(true, nil)
		service := NewPostAdService(repo)

		service.Execute(request)

		repo.AssertCalled(t, "Save", ad)
		assert.NoError(t, nil)
	})

	t.Run("fail when title is larger than 50 characters", func(t *testing.T) {
		repo := NewMockRepository(true)
		service := NewPostAdService(repo)
		request := PostAdRequest{
			"69d39636-d256-47e6-bf86-6bef0cb32ceb",
			"Lorem ipsum dolor sit amet, consectetuer adipiscing",
			"This is a test ad",
			9.99,
			now,
		}

		actualError := service.Execute(request)

		expectedError := fmt.Errorf("%w: %s", ErrTitleTooLong, request.Title())
		assert.Equal(t, expectedError, actualError)
	})

	t.Run("error", func(t *testing.T) {
		repo := NewMockRepository(true)
		expectedError := errors.New("an error")
		repo.On("Save", ad).Return(false, expectedError)
		service := NewPostAdService(repo)

		actualError := service.Execute(request)

		repo.AssertCalled(t, "Save", ad)
		assert.Equal(t, expectedError, actualError)
	})
}