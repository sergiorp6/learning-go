package postgres

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	repo := setupRepository()
	ad, _ := domain.NewAd(
		"78b1c410-3ea7-4e9b-8a4b-9809b5c43394",
		"A title",
		"A description",
		9.99,
		time.Now(),
	)

	saved, err := repo.Save(ad)
	assert.True(t, saved, "Ad should have been saved successfully")
	assert.NoError(t, err)

	found, err := repo.FindBy(ad.Id())
	assert.NoError(t, err)
	assert.True(t, adsAreEqual(ad, *found))

	ads, _ := repo.FindSetOf(5)
	assert.Len(t, ads, 1, "FindSetOf should return one ad")
}

func adsAreEqual(expected domain.Ad, actual domain.Ad) bool {
	return expected.String() == actual.String()
}

func setupRepository() *Repository {
	ctx := context.Background()
	req, pgHost, pgPort := buildDatabaseContainer(ctx)
	db := connectToDatabase(req, pgHost, pgPort)

	return NewRepository(db, ctx)
}

func buildDatabaseContainer(ctx context.Context) (testcontainers.ContainerRequest, string, nat.Port) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14.5",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "123123",
			"POSTGRES_DB":       "test",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not start PostgreSQL container: %s", err)
	}

	pgHost, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get PostgreSQL host: %s", err)
	}

	pgPort, err := container.MappedPort(ctx, "5432/tcp")
	if err != nil {
		log.Fatalf("Could not get PostgreSQL port: %s", err)
	}

	return req, pgHost, pgPort
}

func connectToDatabase(req testcontainers.ContainerRequest, pgHost string, pgPort nat.Port) *gorm.DB {
	pgUsername, _ := req.Env["POSTGRES_USER"]
	pgPassword, _ := req.Env["POSTGRES_PASSWORD"]
	pgDatabase, _ := req.Env["POSTGRES_DB"]

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", pgUsername, pgPassword, pgHost, pgPort.Int(), pgDatabase)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db
}
