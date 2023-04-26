package postad

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"
	"testing"
	"time"
)

type FrozenClock struct {
}

func (f FrozenClock) Now() time.Time {
	input := "2023-01-01"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)

	return t
}

func TestPostAdService_Execute(t *testing.T) {
	clock := FrozenClock{}
	ad, _ := NewAd(
		"69d39636-d256-47e6-bf86-6bef0cb32ceb",
		"Test Ad",
		"This is a test ad",
		9.99,
		clock.Now(),
	)
	request := Request{
		"69d39636-d256-47e6-bf86-6bef0cb32ceb",
		"Test Ad",
		"This is a test ad",
		9.99,
	}

	t.Run("success", func(t *testing.T) {
		repo := NewMockRepository(true)
		repo.On("Save", ad).Return(true, nil)
		service := NewService(repo, clock, dummyEventBus{})

		service.Execute(request)

		repo.AssertCalled(t, "Save", ad)
		assert.NoError(t, nil)
	})

	t.Run("fail when title is larger than 50 characters", func(t *testing.T) {
		repo := NewMockRepository(true)
		service := NewService(repo, clock, dummyEventBus{})
		request := Request{
			"69d39636-d256-47e6-bf86-6bef0cb32ceb",
			"Lorem ipsum dolor sit amet, consectetuer adipiscing",
			"This is a test ad",
			9.99,
		}

		actualError := service.Execute(request)

		expectedError := fmt.Errorf("%w: %s", ErrTitleTooLong, request.Title())
		assert.Equal(t, expectedError, actualError)
	})

	t.Run("error", func(t *testing.T) {
		repo := NewMockRepository(true)
		expectedError := errors.New("an error")
		repo.On("Save", ad).Return(false, expectedError)
		service := NewService(repo, clock, dummyEventBus{})

		actualError := service.Execute(request)

		repo.AssertCalled(t, "Save", ad)
		assert.Equal(t, expectedError, actualError)
	})
}

type dummyEventBus struct {
}

func (d dummyEventBus) Publish(event Event) error {
	return nil
}
