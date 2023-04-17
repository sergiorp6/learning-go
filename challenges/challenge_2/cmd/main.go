package main

import (
	"fmt"
	"github.com/google/uuid"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/application/findbyid"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/application/getadslist"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/application/postad"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_2/internal/ad/infrastructure"
	"time"
)

var repository = &InMemoryRepository{}

func main() {
	for range [10]struct{}{} {
		postAd()
	}
	printAllAds()
}

func postAd() {
	service := NewPostAdService(repository)
	uuidString := uuid.New().String()

	err := service.Execute(
		NewPostAdRequest(
			uuidString,
			"title",
			"description",
			10.50,
			time.Now(),
		),
	)
	if err != nil {
		_ = fmt.Errorf("error posting ad: %s", uuidString)
	}
	printAdContents(uuidString)
}

func printAdContents(id string) {
	service := NewFindAdByIdService(repository)

	ad, err := service.Execute(NewFindAdByIdRequest(id))

	if err != nil {
		_ = fmt.Errorf("error fetching ad: %s", id)
	}

	if ad != nil {
		fmt.Printf("Found ad by id %s: %s\n", id, ad)
	} else {
		fmt.Printf("Ad with id %s not found\n", id)
	}
}

func printAllAds() {
	service := NewGetAdsListService(repository)

	ads := service.Execute(NewGetAdsListRequest(10))

	fmt.Printf("Printing a list of %d ads\n", len(ads))
	for _, ad := range ads {
		fmt.Println(ad)
	}
}
