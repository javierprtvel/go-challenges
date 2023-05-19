package client

import (
	"github.com/gin-gonic/gin"
	"github.mpi-internal.com/javier-porto/learning-go/application"
	"time"
)

type AdController struct {
	router    *gin.Engine
	adService application.AdService
}

func NewAdController(adService application.AdService) AdController {
	router := gin.Default()
	router.GET("/ads/:adId", handleGetAd(adService))
	router.GET("/ads", handleGetSomeAds(adService))
	router.POST("/ads", handleCreateAd(adService))

	return AdController{
		router:    router,
		adService: adService,
	}
}

func (adController AdController) Init() {
	adController.router.Run("localhost:8080")
}

func handleGetAd(adService application.AdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		adId := context.Param("adId")
		getAdResponse := adService.GetAd(application.GetAdRequest{Id: adId})
		httpAdResponse := mapGetAdResponseToHttpResponse(getAdResponse)
		context.JSON(200, httpAdResponse)
	}
}

func handleGetSomeAds(adService application.AdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		getSomeAdsResponse := adService.GetSomeAds()
		httpSomeAdsResponse := mapGetSomeAdsResponseToHttpResponse(getSomeAdsResponse)
		context.JSON(200, httpSomeAdsResponse)
	}
}

func handleCreateAd(adService application.AdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var httpCreateAdRequest httpCreateAdRequest
		context.BindJSON(&httpCreateAdRequest)
		createAdRequest := mapHttpCreateAdRequestToServiceRequest(httpCreateAdRequest)
		adService.CreateAd(createAdRequest)
		context.Status(201)
	}
}

type httpAdResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       uint32    `json:"price"`
	Date        time.Time `json:"date"`
}

func mapGetAdResponseToHttpResponse(gar application.GetAdResponse) httpAdResponse {
	return httpAdResponse{
		Id:          gar.Id,
		Title:       gar.Title,
		Description: gar.Description,
		Price:       gar.Price,
		Date:        gar.Date,
	}
}

type httpAdsResponse struct {
	Ads []httpAdResponse `json:"ads"`
}

func mapGetSomeAdsResponseToHttpResponse(gsar application.GetSomeAdsResponse) httpAdsResponse {
	ads := make([]httpAdResponse, 0)
	for _, gar := range gsar.Ads {
		ads = append(ads, mapGetAdResponseToHttpResponse(gar))
	}
	return httpAdsResponse{ads}
}

type httpCreateAdRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
}

func mapHttpCreateAdRequestToServiceRequest(hcar httpCreateAdRequest) application.CreateAdRequest {
	return application.CreateAdRequest{
		Title:       hcar.Title,
		Description: hcar.Description,
		Price:       hcar.Price,
	}
}
