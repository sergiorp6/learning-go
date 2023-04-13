package application

import (
	"fmt"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_1/internal/ad/domain"
)

type postAdService struct {
	adRepository Repository
}

func NewPostAdService(adRepository Repository) postAdService {
	return postAdService{adRepository}
}

func (s postAdService) Execute(request PostAdRequest) {
	_, err := s.adRepository.Save(
		NewAd(
			NewId(request.Id),
			Title{Value: request.Title},
			Description{Value: request.Description},
			Price{Value: request.Price},
		),
	)
	if err != nil {
		_ = fmt.Errorf("error posting ad: %s")
	}
}

type PostAdRequest struct {
	Id          string
	Title       string
	Description string
	Price       float64
}
