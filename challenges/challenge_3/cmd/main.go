package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/application/findbyid"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/application/getadslist"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/application/postad"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/domain"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/infrastructure/ad"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/infrastructure/server/handler/findbyid"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/infrastructure/server/handler/getadslist"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_3/internal/ad/infrastructure/server/handler/postad"
	"log"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	repository := &InMemoryRepository{}

	srv := gin.Default()
	srv.PUT("/ads/:id", PostAdHandler(NewPostAdService(repository, DefaultClock{})))
	srv.GET("/ads/:id", FindByIdHandler(NewFindByIdService(repository)))
	srv.GET("/ads", GetAdsListHandler(NewGetAdsListService(repository)))

	log.Fatal(srv.Run(httpAddr))
}
