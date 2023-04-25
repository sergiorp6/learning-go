package getadslist

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
	"testing"
	"time"
)

func Test_getAdsListService_Execute(t *testing.T) {
	const numberOfElements = 10
	now := time.Now()
	var expected []Ad

	t.Run("success", func(t *testing.T) {
		for range [numberOfElements]struct{}{} {
			ad, _ := NewAd(
				uuid.NewString(),
				"A title",
				"A description",
				9.99,
				now,
			)
			expected = append(expected, ad)
		}
		repo := NewMockRepository(true)
		repo.On("FindSetOf", numberOfElements).Return(expected, nil)
		service := NewService(repo)

		actual := service.Execute(NewRequest(numberOfElements))

		assert.Equal(t, expected, actual)
	})

	t.Run("empty", func(t *testing.T) {
		expected = []Ad{}
		repo := NewMockRepository(false)
		repo.On("FindSetOf", numberOfElements).Return(expected, nil)
		service := NewService(repo)

		actual := service.Execute(NewRequest(numberOfElements))

		assert.Equal(t, expected, actual)
	})
}
