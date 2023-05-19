package application

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/javier-porto/learning-go/domain"
	"github.mpi-internal.com/javier-porto/learning-go/domain/repository"
	"testing"
	"time"
)

type MockAd domain.Ad

func (ma MockAd) Matches(x interface{}) bool {
	ad := x.(domain.Ad)
	return ad.Title == ma.Title && ad.Description == ma.Description && ad.Price == ma.Price
}
func (ma MockAd) String() string {
	return fmt.Sprintf("{%s %s %d}", ma.Title, ma.Description, ma.Price)
}

func TestAdService_CreateAd(t *testing.T) {
	ctrl := gomock.NewController(t)

	adRepository := repository.NewMockAdRepository(ctrl)
	adService := AdService{adRepository}

	request := CreateAdRequest{
		Title:       "Title for Mock Test",
		Description: "Description Anything",
		Price:       99,
	}

	mockAd := MockAd{
		Title:       "Title for Mock Test",
		Description: "Description Anything",
		Price:       99,
		Date:        time.Now(),
	}
	adRepository.
		EXPECT().
		Persist(gomock.All(mockAd)).
		Return(
			domain.Ad{
				Id:          "15467",
				Title:       "Title for Mock Test",
				Description: "Description Anything",
				Price:       99,
				Date:        time.Now(),
			})

	adService.CreateAd(request)
}

func TestAdService_GetAd(t *testing.T) {
	ctrl := gomock.NewController(t)

	adRepository := repository.NewMockAdRepository(ctrl)
	adService := AdService{adRepository}
	now := time.Now()

	request := GetAdRequest{Id: "4"}

	adRepository.
		EXPECT().
		FindById("4").
		Return(
			domain.Ad{
				Id:          "4",
				Title:       "Title for Mock Test",
				Description: "Description of fourth ad",
				Price:       15,
				Date:        now,
			})

	actual := adService.GetAd(request)

	expected := GetAdResponse{
		Id:          "4",
		Title:       "Title for Mock Test",
		Description: "Description of fourth ad",
		Price:       15,
		Date:        now,
	}
	assert.Equal(t, expected, actual)
}

func TestAdService_GetSomeAds(t *testing.T) {
	ctrl := gomock.NewController(t)

	adRepository := repository.NewMockAdRepository(ctrl)
	adService := AdService{adRepository}
	now := time.Now()

	adRepository.
		EXPECT().
		Slice().
		Return(
			[]domain.Ad{
				{
					Id:          "1",
					Title:       "Title 1",
					Description: "Desc 1",
					Price:       99,
					Date:        now,
				},
				{
					Id:          "2",
					Title:       "Title 2",
					Description: "Desc 2",
					Price:       15,
					Date:        now,
				},
				{
					Id:          "3",
					Title:       "Title 3",
					Description: "Lorem ipsum dolor...",
					Price:       2,
					Date:        now,
				},
			})

	actual := adService.GetSomeAds()

	expected := GetSomeAdsResponse{Ads: []GetAdResponse{
		{
			Id:          "1",
			Title:       "Title 1",
			Description: "Desc 1",
			Price:       99,
			Date:        now,
		},
		{
			Id:          "2",
			Title:       "Title 2",
			Description: "Desc 2",
			Price:       15,
			Date:        now,
		},
		{
			Id:          "3",
			Title:       "Title 3",
			Description: "Lorem ipsum dolor...",
			Price:       2,
			Date:        now,
		},
	}}

	assert.NotSame(t, expected.Ads, actual.Ads)
	assert.Equal(t, expected.Ads, actual.Ads)
}
