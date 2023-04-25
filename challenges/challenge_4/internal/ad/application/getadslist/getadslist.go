package getadslist

import (
	"fmt"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
)

type ServiceInterface interface {
	Execute(request Request) []Ad
}

type Service struct {
	adRepository Repository
}

func NewService(adRepository Repository) Service {
	return Service{adRepository: adRepository}
}

func (s Service) Execute(request Request) []Ad {
	ads, err := s.adRepository.FindSetOf(request.NumberOfElements())
	if err != nil {
		_ = fmt.Errorf("error getting a list of %d ads", request.NumberOfElements())
	}
	return ads
}

type Request struct {
	numberOfElements int
}

func NewRequest(numberOfElements int) Request {
	return Request{numberOfElements: numberOfElements}
}

func (g Request) NumberOfElements() int {
	return g.numberOfElements
}
