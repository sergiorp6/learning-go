package findbyid

import (
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/domain"
)

type FindByIdServiceInterface interface {
	Execute(request FindByIdRequest) (*Ad, error)
}

type FindByIdService struct {
	adRepository Repository
}

func NewFindByIdService(adRepository Repository) FindByIdService {
	return FindByIdService{adRepository: adRepository}
}

func (s FindByIdService) Execute(request FindByIdRequest) (*Ad, error) {
	return s.adRepository.FindBy(NewId(request.Id()))
}

type FindByIdRequest struct {
	id string
}

func NewFindByIdRequest(id string) FindByIdRequest {
	return FindByIdRequest{id: id}
}

func (f FindByIdRequest) Id() string {
	return f.id
}
