package application

import (
	"fmt"
	"github.com/google/uuid"
	"github.mpi-internal.com/javier-porto/learning-go/domain"
	"github.mpi-internal.com/javier-porto/learning-go/domain/repository"
	"time"
)

type CreateAdRequest struct {
	Title, Description string
	Price              uint32
}

type GetAdRequest struct {
	Id string
}
type GetAdResponse struct {
	Id, Title, Description string
	Price                  uint32
	Date                   time.Time
}

type GetSomeAdsResponse struct {
	Ads []GetAdResponse
}

type AdService struct {
	repository repository.AdRepository
}

func NewAdService(adRepository repository.AdRepository) AdService {
	return AdService{adRepository}
}

func (adService AdService) CreateAd(request CreateAdRequest) error {
	title := request.Title
	adsWithSameTitle := adService.repository.FindByTitle(title)
	if len(adsWithSameTitle) > 0 {
		return fmt.Errorf("an ad with title %s already exists", title)
	}
	adService.repository.Persist(domain.Ad{
		Id:          uuid.NewString(),
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Date:        time.Now(),
	})
	return nil
}

func (adService AdService) GetAd(request GetAdRequest) GetAdResponse {
	ad := adService.repository.FindById(request.Id)
	return GetAdResponse{
		Id:          ad.Id,
		Title:       ad.Title,
		Description: ad.Description,
		Price:       ad.Price,
		Date:        ad.Date,
	}
}

func (adService AdService) GetSomeAds() GetSomeAdsResponse {
	ads := adService.repository.Slice()
	getSomeAdsResponse := GetSomeAdsResponse{}
	for _, ad := range ads {
		getSomeAdsResponse.Ads = append(getSomeAdsResponse.Ads, GetAdResponse{
			Id:          ad.Id,
			Title:       ad.Title,
			Description: ad.Description,
			Price:       ad.Price,
			Date:        ad.Date,
		})
	}
	return getSomeAdsResponse
}
