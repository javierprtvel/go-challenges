package client

import (
	"fmt"
	"github.mpi-internal.com/javier-porto/learning-go/application"
)

type CLI struct{}

func (cli CLI) AskUserForAdId() string {
	var adId string
	fmt.Println("Search ad by id:")
	fmt.Scan(&adId)
	return adId
}

func (cli CLI) ShowAdToUser(ad application.GetAdResponse) {
	if &ad != nil {
		fmt.Println("Ad:", ad)
	} else {
		fmt.Println("Ad not found")
	}
}
