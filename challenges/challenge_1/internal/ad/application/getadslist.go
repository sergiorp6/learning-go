package application

import (
	"fmt"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_1/internal/ad/domain"
)

type getAdsListService struct {
	adRepository Repository
}

func NewGetAdsListService(adRepository Repository) getAdsListService {
	return getAdsListService{adRepository: adRepository}
}
func (s getAdsListService) Execute(request GetAdsListRequest) []Ad {
	ads, err := s.adRepository.FindSetOf(request.NumberOfElements)
	if err != nil {
		_ = fmt.Errorf("error getting a list of %d ads", request.NumberOfElements)
	}
	return ads
}

type GetAdsListRequest struct {
	NumberOfElements int
}
