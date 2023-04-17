package findbyid

import (
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
)

type findAdByIdService struct {
	adRepository Repository
}

func NewFindAdByIdService(adRepository Repository) findAdByIdService {
	return findAdByIdService{adRepository: adRepository}
}

func (s findAdByIdService) Execute(request FindAdByIdRequest) (*Ad, error) {
	return s.adRepository.FindBy(NewId(request.Id()))
}

type FindAdByIdRequest struct {
	id string
}

func NewFindAdByIdRequest(id string) FindAdByIdRequest {
	return FindAdByIdRequest{id: id}
}

func (f FindAdByIdRequest) Id() string {
	return f.id
}
