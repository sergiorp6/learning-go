package application

import (
	"fmt"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_1/internal/ad/domain"
)

type findAdByIdService struct {
	adRepository Repository
}

func NewFindAdByIdService(adRepository Repository) findAdByIdService {
	return findAdByIdService{adRepository: adRepository}
}

func (s findAdByIdService) Execute(request FindAdByIdRequest) *Ad {
	ad, err := s.adRepository.FindBy(NewId(request.Id))
	if err != nil {
		_ = fmt.Errorf("error finding ad by id %s", request.Id)
	}
	return ad
}

type FindAdByIdRequest struct {
	Id string
}
