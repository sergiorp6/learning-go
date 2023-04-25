package postgres

import (
	"context"
	"fmt"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewRepository(db *gorm.DB, ctx context.Context) *Repository {
	initDatabase(db)
	return &Repository{db, ctx}
}

func (r *Repository) Save(ad domain.Ad) (bool, error) {
	err := r.db.Create(&Ad{
		Id:          ad.Id().String(),
		Title:       ad.Title().Value(),
		Description: ad.Description().Value(),
		Price:       ad.Price().Value(),
		PublishedAt: ad.PublishedAt().Value().Format(time.RFC3339),
	}).Error
	if err != nil {
		return false, fmt.Errorf("could not save ad: %w", err)
	}
	return true, nil
}

func (r *Repository) FindBy(id domain.Id) (*domain.Ad, error) {
	var ad Ad
	err := r.db.First(&ad, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding ad by id: %w", err)
	}

	domainAd, err := ad.toDomain()

	return domainAd, nil
}

func (r *Repository) FindSetOf(number int) ([]domain.Ad, error) {
	var ads []Ad
	err := r.db.Limit(number).Find(&ads).Error
	if err != nil {
		return nil, fmt.Errorf("could not find set of ads: %w", err)
	}
	var domainAds []domain.Ad

	for _, ad := range ads {
		domainAd, err := ad.toDomain()
		if err != nil {
			return nil, fmt.Errorf("error converting ad to domain entity: %w", err)
		}
		domainAds = append(domainAds, *domainAd)
	}

	return domainAds, nil
}

func initDatabase(db *gorm.DB) {
	err := db.AutoMigrate(Ad{})
	if err != nil {
		log.Fatalf("Error migrating ad: %v", err)
	}
}

type Ad struct {
	Id          string
	Title       string
	Description string
	Price       float64
	PublishedAt string
}

func (a Ad) toDomain() (*domain.Ad, error) {
	publishedAt, err := time.Parse(time.RFC3339, a.PublishedAt)
	if err != nil {
		return nil, fmt.Errorf("error parsing publishedAt: %w", err)
	}

	domainAd, err := domain.NewAd(
		a.Id,
		a.Title,
		a.Description,
		a.Price,
		publishedAt,
	)

	return &domainAd, nil
}
