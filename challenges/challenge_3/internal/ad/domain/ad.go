package domain

import (
	"errors"
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

func NewAd(id, title, description string, price float64, publishedAt time.Time) (Ad, error) {
	titleVo, err := NewTitle(title)

	if err != nil {
		return Ad{}, err
	}

	return Ad{
		id:          NewId(id),
		title:       titleVo,
		description: NewDescription(description),
		price:       NewPrice(price),
		publishedAt: NewPublishedAt(publishedAt),
	}, nil
}

func (ad Ad) String() string {
	return fmt.Sprintf("Ad{ID:%s, Title:%s, Description:%s, Price:%f, PublishedAt:%s}",
		ad.id, ad.title, ad.description, ad.price, ad.publishedAt.Value().Format(time.RFC3339))
}
func (ad Ad) Id() Id {
	return ad.id
}

func (ad Ad) Title() Title {
	return ad.title
}

func (ad Ad) Description() Description {
	return ad.description
}

func (ad Ad) Price() Price {
	return ad.price
}

func (ad Ad) PublishedAt() PublishedAt {
	return ad.publishedAt
}

type Id struct {
	value uuid.UUID
}

func NewId(uuidString string) Id {
	parsed, err := uuid.Parse(uuidString)
	if err != nil {
		_ = fmt.Errorf("%w: %s", err, uuidString)
	}
	return Id{value: parsed}
}

func (id Id) String() string {
	return id.value.String()
}

type Title struct {
	value string
}

var ErrTitleTooLong = errors.New("Title above ")

func NewTitle(value string) (Title, error) {
	if len(value) > 50 {
		return Title{}, fmt.Errorf("%w: %s", ErrTitleTooLong, value)
	}
	return Title{value: value}, nil
}

func (t Title) Value() string {
	return t.value
}

type Description struct {
	value string
}

func NewDescription(value string) Description {
	return Description{value: value}
}

func (d Description) Value() string {
	return d.value
}

type Price struct {
	value float64
}

func NewPrice(value float64) Price {
	return Price{value: value}
}

func (p Price) Value() float64 {
	return p.value
}

type PublishedAt struct {
	value time.Time
}

func NewPublishedAt(value time.Time) PublishedAt {
	return PublishedAt{value: value}
}

func (p PublishedAt) Value() time.Time {
	return p.value
}

type Repository interface {
	Save(ad Ad) (bool, error)
	FindBy(id Id) (*Ad, error)
	FindSetOf(number int) ([]Ad, error)
}
