package main

import (
	"fmt"
	"github.com/google/uuid"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_1/internal/ad/application"
	. "github.mpi-internal.com/sergio.rodriguezp/learning-go/challenges/challenge_1/internal/ad/infrastructure"
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

	service.Execute(PostAdRequest{Id: uuidString, Title: "title", Description: "description", Price: 10.50})
	printAdContents(uuidString)
}

func printAdContents(id string) {
	service := NewFindAdByIdService(repository)

	ad := service.Execute(FindAdByIdRequest{Id: id})

	if ad != nil {
		fmt.Printf("Found ad by id %s: %s\n", id, ad)
	} else {
		fmt.Printf("Ad with id %s not found\n", id)
	}
}

func printAllAds() {
	service := NewGetAdsListService(repository)

	ads := service.Execute(GetAdsListRequest{NumberOfElements: 10})

	fmt.Printf("Printing a list of %d ads\n", len(ads))
	for _, ad := range ads {
		fmt.Println(ad)
	}
}
