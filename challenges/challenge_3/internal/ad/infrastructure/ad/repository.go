package ad

import . "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/domain"

type InMemoryRepository struct {
	data []Ad
}

func (r *InMemoryRepository) Save(ad Ad) (bool, error) {
	r.data = append(r.data, ad)
	return true, nil
}

func (r *InMemoryRepository) FindBy(id Id) (*Ad, error) {
	for i := 0; i < len(r.data); i++ {
		if r.data[i].Id() == id {
			return &r.data[i], nil
		}
	}
	return nil, nil
}

func (r *InMemoryRepository) FindSetOf(number int) ([]Ad, error) {
	const min = 1
	const max = 5

	if len(r.data) == 0 {
		return nil, nil
	}

	if number < min {
		number = min
	} else if number > max {
		number = max
	}

	return r.data[:number], nil
}
