package main

import (
	"github.mpi-internal.com/javier-porto/learning-go/application"
	"github.mpi-internal.com/javier-porto/learning-go/infrastructure/client"
	"github.mpi-internal.com/javier-porto/learning-go/infrastructure/repository"
)

func main() {
	adService := application.NewAdService(repository.NewInMemoryAdRepository())
	setInitialAdCatalog(adService)

	server := client.SetupServer(adService)
	server.Run(":8080")
}

func setInitialAdCatalog(adService application.AdService) {
	adService.CreateAd(application.CreateAdRequest{
		Title:       "Title 1",
		Description: "No description",
		Price:       28,
	})
	adService.CreateAd(application.CreateAdRequest{
		Title:       "Title 2",
		Description: "No description",
		Price:       50,
	})
	adService.CreateAd(application.CreateAdRequest{
		Title:       "Title 3",
		Description: "No description",
		Price:       17,
	})
	adService.CreateAd(application.CreateAdRequest{
		Title:       "Title 4",
		Description: "No description",
		Price:       21,
	})
	adService.CreateAd(application.CreateAdRequest{
		Title:       "Title 5",
		Description: "No description",
		Price:       99,
	})
	adService.CreateAd(application.CreateAdRequest{
		Title:       "Title 6",
		Description: "No description",
		Price:       5,
	})
}
