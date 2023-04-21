package postad

import (
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/domain"
)

type PostAdServiceInterface interface {
	Execute(request PostAdRequest) error
}

type PostAdService struct {
	adRepository Repository
	clock        Clock
}

func NewPostAdService(adRepository Repository, clock Clock) PostAdService {
	return PostAdService{adRepository, clock}
}

func (s PostAdService) Execute(request PostAdRequest) error {
	ad, err := NewAd(
		request.Id(),
		request.Title(),
		request.Description(),
		request.Price(),
		s.clock.Now(),
	)
	if err != nil {
		return err
	}
	_, err = s.adRepository.Save(ad)

	return err
}

type PostAdRequest struct {
	id          string
	title       string
	description string
	price       float64
}

func NewPostAdRequest(id, title, description string, price float64) PostAdRequest {
	return PostAdRequest{id: id, title: title, description: description, price: price}
}

func (p PostAdRequest) Id() string {
	return p.id
}

func (p PostAdRequest) Title() string {
	return p.title
}

func (p PostAdRequest) Description() string {
	return p.description
}

func (p PostAdRequest) Price() float64 {
	return p.price
}
