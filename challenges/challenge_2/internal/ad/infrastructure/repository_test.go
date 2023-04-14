package infrastructure

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
	"testing"
	"time"
)

func TestInMemoryRepository_Save(t *testing.T) {
	repo := InMemoryRepository{}
	ad := NewAd(
		NewId("78b1c410-3ea7-4e9b-8a4b-9809b5c43394"),
		Title{Value: "A title"},
		Description{Value: "A description"},
		Price{Value: 9.99},
		PublishedAt{Value: time.Now()},
	)

	repo.Save(ad)

	assert.Equal(t, 1, len(repo.data))
	assert.Equal(t, ad, repo.data[0])
}

func TestInMemoryRepository_FindBy(t *testing.T) {
	existingId := NewId("78b1c410-3ea7-4e9b-8a4b-9809b5c43394")
	nonExistingId := NewId("b41f9d3a-5036-4d21-8a92-fedfefca11e8")
	expected := NewAd(
		existingId,
		Title{Value: "A title"},
		Description{Value: "A description"},
		Price{Value: 9.99},
		PublishedAt{Value: time.Now()},
	)
	repo := InMemoryRepository{[]Ad{expected}}

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
		repo := InMemoryRepository{}

		actual, _ := repo.FindSetOf(10)

		assert.Empty(t, actual)
	})

	t.Run("Get less elements than maximum allowed", func(t *testing.T) {
		const numberOfElements = 4
		var ads []Ad

		for range [numberOfElements]struct{}{} {
			ads = append(
				ads,
				NewAd(
					NewId(uuid.NewString()),
					Title{Value: "A title"},
					Description{Value: "A description"},
					Price{Value: 9.99},
					PublishedAt{Value: time.Now()},
				),
			)
		}
		repo := InMemoryRepository{ads}

		actual, _ := repo.FindSetOf(numberOfElements)

		assert.Len(t, actual, numberOfElements)
	})

	t.Run("Get the maximum elements allowed", func(t *testing.T) {
		const numberOfElements = 10
		var ads []Ad

		for range [numberOfElements]struct{}{} {
			ads = append(
				ads,
				NewAd(
					NewId(uuid.NewString()),
					Title{Value: "A title"},
					Description{Value: "A description"},
					Price{Value: 9.99},
					PublishedAt{Value: time.Now()},
				),
			)
		}

		repo := InMemoryRepository{ads}
		actual, _ := repo.FindSetOf(numberOfElements)

		const maxNumberOfElements = 5
		assert.Len(t, actual, maxNumberOfElements)
	})
}
