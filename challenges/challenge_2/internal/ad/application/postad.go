package application

import (
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/domain"
	"time"
)

type postAdService struct {
	adRepository Repository
}

func NewPostAdService(adRepository Repository) postAdService {
	return postAdService{adRepository}
}

func (s postAdService) Execute(request PostAdRequest) error {
	_, err := s.adRepository.Save(
		NewAd(
			NewId(request.Id),
			Title{Value: request.Title},
			Description{Value: request.Description},
			Price{Value: request.Price},
			PublishedAt{Value: request.PublishedAt},
		),
	)
	return err
}

type PostAdRequest struct {
	Id          string
	Title       string
	Description string
	Price       float64
	PublishedAt time.Time
}
