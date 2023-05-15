package main

import (
	"fmt"
	"github.mpi-internal.com/javier-porto/learning-go/application"
	"github.mpi-internal.com/javier-porto/learning-go/infrastructure/client"
)

func main() {
	cli := client.CLI{}
	adService := application.InitAdService()

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

	adListing := adService.GetSomeAds()
	fmt.Println("Ad listing:", adListing.Ads)
	fmt.Println("Ad listing size:", len(adListing.Ads))

	adId := cli.AskUserForAdId()
	userAd := adService.GetAd(application.GetAdRequest{Id: adId})
	cli.ShowAdToUser(userAd)
}
