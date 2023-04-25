package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/application/findbyid"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/application/getadslist"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/application/postad"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/domain"
	findbyidhandler "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/infrastructure/handler/findbyid"
	getadslisthandler "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/infrastructure/handler/getadslist"
	postadhandler "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/infrastructure/handler/postad"
	repository "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_4/internal/ad/infrastructure/repository/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const httpAddr = ":8080"

func main() {
	startupWebServer()
}

func startupWebServer() {
	fmt.Println("Server running on", httpAddr)
	repo := initRepository()
	srv := gin.Default()
	setUpRoutes(srv, repo)

	log.Fatal(srv.Run(httpAddr))
}

func setUpRoutes(srv *gin.Engine, repository *repository.Repository) {
	srv.PUT("/ads/:id", postadhandler.Handler(postad.NewService(repository, domain.DefaultClock{})))
	srv.GET("/ads/:id", findbyidhandler.Handler(findbyid.NewService(repository)))
	srv.GET("/ads", getadslisthandler.Handler(getadslist.NewService(repository)))
}

func initRepository() *repository.Repository {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return repository.NewRepository(db, context.Background())
}
