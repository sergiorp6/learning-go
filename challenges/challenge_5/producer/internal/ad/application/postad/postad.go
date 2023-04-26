package postad

import (
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/application"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"
)

type ServiceInterface interface {
	Execute(request Request) error
}

type Service struct {
	adRepository Repository
	clock        Clock
	eventBus     application.EventBus
}

func NewService(adRepository Repository, clock Clock, bus application.EventBus) Service {
	return Service{adRepository, clock, bus}
}

func (s Service) Execute(request Request) error {
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

	err = s.eventBus.Publish(NewAdPosted(ad))

	if err != nil {
		return err
	}

	return err
}

type Request struct {
	id          string
	title       string
	description string
	price       float64
}

func NewRequest(id, title, description string, price float64) Request {
	return Request{id: id, title: title, description: description, price: price}
}

func (p Request) Id() string {
	return p.id
}

func (p Request) Title() string {
	return p.title
}

func (p Request) Description() string {
	return p.description
}

func (p Request) Price() float64 {
	return p.price
}
