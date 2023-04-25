package inmemory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
	"testing"
	"time"
)

func TestInMemoryRepository_Save(t *testing.T) {
	repo := Repository{}
	ad, _ := domain.NewAd(
		"78b1c410-3ea7-4e9b-8a4b-9809b5c43394",
		"A title",
		"A description",
		9.99,
		time.Now(),
	)

	repo.Save(ad)

	assert.Equal(t, 1, len(repo.data))
	assert.Equal(t, ad, repo.data[0])
}

func TestInMemoryRepository_FindBy(t *testing.T) {
	existingId := domain.NewId("78b1c410-3ea7-4e9b-8a4b-9809b5c43394")
	nonExistingId := domain.NewId("b41f9d3a-5036-4d21-8a92-fedfefca11e8")
	expected, _ := domain.NewAd(
		existingId.String(),
		"A title",
		"A description",
		9.99,
		time.Now(),
	)
	repo := Repository{[]domain.Ad{expected}}

	t.Run("Existing ad", func(t *testing.T) {
		actual, _ := repo.FindBy(existingId)

		assert.Equal(t, &expected, actual)
	})

	t.Run("Non existing ad", func(t *testing.T) {
		actual, _ := repo.FindBy(nonExistingId)

		assert.Nil(t, actual)
	})
}

func TestInMemoryRepository_FindSetOf(t *testing.T) {
	t.Run("Find empty list", func(t *testing.T) {
		repo := Repository{}

		actual, _ := repo.FindSetOf(10)

		assert.Empty(t, actual)
	})

	t.Run("Get less elements than maximum allowed", func(t *testing.T) {
		const numberOfElements = 4
		var ads []domain.Ad
		for range [numberOfElements]struct{}{} {
			ad, _ := domain.NewAd(
				uuid.NewString(),
				"A title",
				"A description",
				9.99,
				time.Now(),
			)
			ads = append(ads, ad)
		}
		repo := Repository{ads}

		actual, _ := repo.FindSetOf(numberOfElements)

		assert.Len(t, actual, numberOfElements)
	})

	t.Run("Get the maximum elements allowed", func(t *testing.T) {
		const numberOfElements = 10
		var ads []domain.Ad
		for range [numberOfElements]struct{}{} {
			ad, _ := domain.NewAd(
				uuid.NewString(),
				"A title",
				"A description",
				9.99,
				time.Now(),
			)
			ads = append(ads, ad)
		}
		repo := Repository{ads}

		actual, _ := repo.FindSetOf(numberOfElements)

		const maxNumberOfElements = 5
		assert.Len(t, actual, maxNumberOfElements)
	})
	t.Run("Try to get more elements than persisted", func(t *testing.T) {
		const numberOfElements = 10
		const persistedElements = 1
		var ads []domain.Ad
		for range [persistedElements]struct{}{} {
			ad, _ := domain.NewAd(
				uuid.NewString(),
				"A title",
				"A description",
				9.99,
				time.Now(),
			)
			ads = append(ads, ad)
		}
		repo := Repository{ads}

		actual, _ := repo.FindSetOf(numberOfElements)

		assert.Len(t, actual, persistedElements)
	})

}
