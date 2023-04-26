package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/application/findbyid"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/application/getadslist"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/application/postad"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/domain"
	"github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/infrastructure/eventbus/kafka"
	findbyidhandler "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/infrastructure/handler/findbyid"
	getadslisthandler "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/infrastructure/handler/getadslist"
	postadhandler "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/infrastructure/handler/postad"
	repository "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_5/producer/internal/ad/infrastructure/repository/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

const httpAddr = ":8080"

func main() {
	createKafkaTopic()
	startupWebServer()
}

func createKafkaTopic() {
	brokerAddrs := []string{os.Getenv("KAFKA_BROKER")}
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)

	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() { _ = admin.Close() }()
	err = admin.CreateTopic(os.Getenv("KAFKA_TOPIC"), &sarama.TopicDetail{
		NumPartitions:     4,
		ReplicationFactor: 1,
	}, false)
	if err != nil {
		if !strings.Contains(err.Error(), "Topic with this name already exists") {
			log.Fatal("Error while creating topic: ", err.Error())
		}
	}
}

func startupWebServer() {
	fmt.Println("Server running on", httpAddr)
	eventBus := initEventBus()
	repo := initRepository()
	srv := gin.Default()
	setUpRoutes(srv, repo, eventBus)

	log.Fatal(srv.Run(httpAddr))
}

func setUpRoutes(srv *gin.Engine, repository *repository.Repository, bus *kafka.EventBus) {
	srv.PUT("/ads/:id", postadhandler.Handler(postad.NewService(repository, domain.DefaultClock{}, bus)))
	srv.GET("/ads/:id", findbyidhandler.Handler(findbyid.NewService(repository)))
	srv.GET("/ads", getadslisthandler.Handler(getadslist.NewService(repository)))
}

func initEventBus() *kafka.EventBus {
	return kafka.NewEventBus(os.Getenv("KAFKA_BROKER"), os.Getenv("KAFKA_TOPIC"))
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
