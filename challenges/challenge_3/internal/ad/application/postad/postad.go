package postad

import (
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/domain"
	"time"
)

type PostAdService struct {
	adRepository Repository
}

func NewPostAdService(adRepository Repository) PostAdService {
	return PostAdService{adRepository}
}

func (s PostAdService) Execute(request PostAdRequest) error {
	ad, err := NewAd(
		request.Id(),
		request.Title(),
		request.Description(),
		request.Price(),
		request.PublishedAt(),
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
	publishedAt time.Time
}

func NewPostAdRequest(id, title, description string, price float64, publishedAt time.Time) PostAdRequest {
	return PostAdRequest{id: id, title: title, description: description, price: price, publishedAt: publishedAt}
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
func (p PostAdRequest) PublishedAt() time.Time {
	return p.publishedAt
}
