package domain

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	id          Id
	title       Title
	description Description
	price       Price
	publishedAt PublishedAt
}

func NewAd(id Id, title Title, description Description, price Price, publishedAt PublishedAt) Ad {
	return Ad{
		id:          id,
		title:       title,
		description: description,
		price:       price,
		publishedAt: publishedAt,
	}
}

func (ad Ad) String() string {
	return fmt.Sprintf("Ad{ID:%s, Title:%s, Description:%s, Price:%f, PublishedAt:%s}",
		ad.id, ad.title, ad.description, ad.price, ad.publishedAt.Value.Format(time.RFC3339))
}
func (ad Ad) Id() Id {
	return ad.id
}

type Id struct {
	Value uuid.UUID
}

func NewId(uuidString string) Id {
	parsed, err := uuid.Parse(uuidString)
	if err != nil {
		_ = fmt.Errorf("invalid UUID string", err)
	}
	return Id{Value: parsed}
}

func (id Id) String() string {
	return id.Value.String()
}

type Title struct {
	Value string
}

type Description struct {
	Value string
}

type Price struct {
	Value float64
}

type PublishedAt struct {
	Value time.Time
}

type Repository interface {
	Save(ad Ad) (bool, error)
	FindBy(id Id) (*Ad, error)
	FindSetOf(number int) ([]Ad, error)
}
