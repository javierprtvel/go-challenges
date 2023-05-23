package client

import (
	"github.com/gin-gonic/gin"
	"github.mpi-internal.com/javier-porto/learning-go/application"
	"time"
)

func SetupServer(adService application.AdService) *gin.Engine {
	router := gin.Default()
	router.GET("/ads/:adId", handleGetAd(adService))
	router.GET("/ads", handleGetSomeAds(adService))
	router.POST("/ads", handleCreateAd(adService))
	return router
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
		var httpCreateAdRequest HttpCreateAdRequest
		err := context.BindJSON(&httpCreateAdRequest)
		if err != nil {
			return
		}

		createAdRequest := mapHttpCreateAdRequestToServiceRequest(httpCreateAdRequest)
		if err := adService.CreateAd(createAdRequest); err != nil {
			context.JSON(409, gin.H{
				"code":  409,
				"title": "ad-already-exists",
				"error": err.Error(),
			})
		}
		context.Status(201)
	}
}

type HttpAdResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       uint32    `json:"price"`
	Date        time.Time `json:"date"`
}

type HttpAdsResponse struct {
	Ads []HttpAdResponse `json:"ads"`
}

type HttpCreateAdRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       uint32 `json:"price" binding:"required"`
}

type HttpErrorResponse struct {
	Code  uint32 `json:"code"`
	Title string `json:"title"`
	Error string `json:"error"`
}

func mapGetAdResponseToHttpResponse(gar application.GetAdResponse) HttpAdResponse {
	return HttpAdResponse{
		Id:          gar.Id,
		Title:       gar.Title,
		Description: gar.Description,
		Price:       gar.Price,
		Date:        gar.Date,
	}
}

func mapGetSomeAdsResponseToHttpResponse(gsar application.GetSomeAdsResponse) HttpAdsResponse {
	ads := make([]HttpAdResponse, 0)
	for _, gar := range gsar.Ads {
		ads = append(ads, mapGetAdResponseToHttpResponse(gar))
	}
	return HttpAdsResponse{ads}
}

func mapHttpCreateAdRequestToServiceRequest(hcar HttpCreateAdRequest) application.CreateAdRequest {
	return application.CreateAdRequest{
		Title:       hcar.Title,
		Description: hcar.Description,
		Price:       hcar.Price,
	}
}
