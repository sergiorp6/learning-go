package application

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
	"testing"
	"time"
)

func Test_getAdsListService_Execute(t *testing.T) {
	const numberOfElements = 10
	now := time.Now()
	var expected []Ad

	t.Run("success", func(t *testing.T) {
		for range [numberOfElements]struct{}{} {
			expected = append(
				expected,
				NewAd(
					NewId(uuid.NewString()),
					Title{Value: "A title"},
					Description{Value: "A description"},
					Price{Value: 9.99},
					PublishedAt{Value: now},
				),
			)
		}
		repo := NewMockRepository(true)
		repo.On("FindSetOf", numberOfElements).Return(expected, nil)
		service := NewGetAdsListService(repo)

		actual := service.Execute(GetAdsListRequest{NumberOfElements: numberOfElements})

		assert.Equal(t, expected, actual)
	})

	t.Run("empty", func(t *testing.T) {
		expected = []Ad{}
		repo := NewMockRepository(false)
		repo.On("FindSetOf", numberOfElements).Return(expected, nil)
		service := NewGetAdsListService(repo)

		actual := service.Execute(GetAdsListRequest{NumberOfElements: numberOfElements})

		assert.Equal(t, expected, actual)
	})
}
