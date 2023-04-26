package inmemory

import "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"

type Repository struct {
	data []domain.Ad
}

func (r *Repository) Save(ad domain.Ad) (bool, error) {
	r.data = append(r.data, ad)
	return true, nil
}

func (r *Repository) FindBy(id domain.Id) (*domain.Ad, error) {
	for i := 0; i < len(r.data); i++ {
		if r.data[i].Id() == id {
			return &r.data[i], nil
		}
	}
	return nil, nil
}

func (r *Repository) FindSetOf(number int) ([]domain.Ad, error) {
	const min = 1
	const max = 5

	if len(r.data) == 0 {
		return nil, nil
	}

	if len(r.data) < number {
		number = len(r.data)
	}

	if number < min {
		number = min
	} else if number > max {
		number = max
	}

	return r.data[:number], nil
}
