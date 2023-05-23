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

	t.Run("Persist ad in the repository", func(t *testing.T) {
		request := CreateAdRequest{
			Title:       "Title for Mock Test",
			Description: "Description Anything",
			Price:       99,
		}

		adRepository.
			EXPECT().
			FindByTitle(request.Title).
			Return([]domain.Ad{})
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

		err := adService.CreateAd(request)
		assert.Nil(t, err)
	})

	t.Run("Return an error if an ad with same title already exists", func(t *testing.T) {
		request := CreateAdRequest{
			Title:       "Title for Mock Test",
			Description: "Description Anything",
			Price:       99,
		}

		date, _ := time.Parse("2023-03-03 14:00:00.000", "2023-03-03 14:00:00.000")
		adRepository.
			EXPECT().
			FindByTitle(request.Title).
			Return([]domain.Ad{
				{
					Id:          "5dd980ae-2b31-4442-a010-6a7808b0961f",
					Title:       "Title for Mock Test",
					Description: "Another description",
					Price:       56,
					Date:        date,
				},
			})

		err := adService.CreateAd(request)
		assert.NotNil(t, err)
	})

	t.Run("Return an error if the ad data is invalid", func(t *testing.T) {
		request := CreateAdRequest{
			Title:       "Title for Mock Test",
			Description: "Lorem ipsum dolor sit aemet this description is so long that it won't fit in the database",
			Price:       4449,
		}

		adRepository.
			EXPECT().
			FindByTitle(request.Title).
			Return([]domain.Ad{})

		err := adService.CreateAd(request)
		assert.NotNil(t, err)
	})
}

func TestAdService_GetAd(t *testing.T) {
	ctrl := gomock.NewController(t)

	adRepository := repository.NewMockAdRepository(ctrl)
	adService := AdService{adRepository}
	now := time.Now()

	t.Run("Retrieve an ad from the repository", func(t *testing.T) {
		request := GetAdRequest{Id: "4"}

		adRepository.
			EXPECT().
			FindById("4").
			Return(
				&domain.Ad{
					Id:          "4",
					Title:       "Title for Mock Test",
					Description: "Description of fourth ad",
					Price:       15,
					Date:        now,
				})

		actual := adService.GetAd(request)

		expected := &GetAdResponse{
			Id:          "4",
			Title:       "Title for Mock Test",
			Description: "Description of fourth ad",
			Price:       15,
			Date:        now,
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("Return nil if ad does not exist in the repository", func(t *testing.T) {
		request := GetAdRequest{Id: "9"}

		adRepository.
			EXPECT().
			FindById("9").
			Return(nil)

		actual := adService.GetAd(request)

		assert.Nil(t, actual)
	})
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
