package findbyid

import (
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
)

type ServiceInterface interface {
	Execute(request Request) (*Ad, error)
}

type Service struct {
	adRepository Repository
}

func NewService(adRepository Repository) Service {
	return Service{adRepository: adRepository}
}

func (s Service) Execute(request Request) (*Ad, error) {
	return s.adRepository.FindBy(NewId(request.Id()))
}

type Request struct {
	id string
}

func NewRequest(id string) Request {
	return Request{id: id}
}

func (f Request) Id() string {
	return f.id
}
