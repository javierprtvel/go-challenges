package repository

import (
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/javier-porto/learning-go/domain"
	"testing"
	"time"
)

func TestInMemoryAdRepository_Persist(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name     string
		input    domain.Ad
		expected map[string]domain.Ad
	}{
		{
			"Happy Path",
			domain.Ad{
				Id:          "1",
				Title:       "Title 1",
				Description: "Description 1",
				Price:       50,
				Date:        now,
			},
			map[string]domain.Ad{
				"1": {
					Id:          "1",
					Title:       "Title 1",
					Description: "Description 1",
					Price:       50,
					Date:        now,
				},
			},
		},
		{
			"No Description",
			domain.Ad{
				Id:          "2",
				Title:       "Title 2",
				Description: "",
				Price:       25,
				Date:        now,
			},
			map[string]domain.Ad{
				"2": {
					Id:          "2",
					Title:       "Title 2",
					Description: "",
					Price:       25,
					Date:        now,
				},
			},
		},
		{
			"No Date Provided",
			domain.Ad{
				Id:          "3",
				Title:       "Title 3",
				Description: "Description 3",
				Price:       19,
				Date:        time.Time{},
			},
			map[string]domain.Ad{
				"3": {
					Id:          "3",
					Title:       "Title 3",
					Description: "Description 3",
					Price:       19,
					Date:        now,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inMemoryAdRepository := inMemoryAdRepository{}

			returnedValue := inMemoryAdRepository.Persist(tc.input)

			assert.Len(t, inMemoryAdRepository, len(tc.expected))
			actualValue := inMemoryAdRepository[tc.input.Id]
			expectedValue := tc.expected[tc.input.Id]
			assert.Condition(t, compareAds(actualValue, expectedValue))
			assert.Condition(t, compareAds(returnedValue, expectedValue))
		})
	}
}

func TestInMemoryAdRepository_FindById(t *testing.T) {
	type findByInput struct {
		id          string
		existingAds map[string]domain.Ad
	}
	now := time.Now()
	existingAds := map[string]domain.Ad{
		"1": {
			Id:          "1",
			Title:       "Title 1",
			Description: "Description 1",
			Price:       50,
			Date:        now,
		},
		"2": {
			Id:          "2",
			Title:       "Title 2",
			Description: "Description 2",
			Price:       19,
			Date:        now,
		},
		"3": {
			Id:          "3",
			Title:       "Title 3",
			Description: "Description 3",
			Price:       45,
			Date:        now,
		},
	}
	testCases := []struct {
		name     string
		input    findByInput
		expected domain.Ad
	}{
		{
			name: "Happy Path",
			input: findByInput{
				id:          "1",
				existingAds: existingAds,
			},
			expected: domain.Ad{
				Id:          "1",
				Title:       "Title 1",
				Description: "Description 1",
				Price:       50,
				Date:        now,
			},
		},
		{
			name: "Not Found",
			input: findByInput{
				id:          "6",
				existingAds: existingAds,
			},
			expected: domain.Ad{
				Id:          "",
				Title:       "",
				Description: "",
				Price:       0,
				Date:        time.Time{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var inMemoryAdRepository inMemoryAdRepository = existingAds

			actual := inMemoryAdRepository.FindById(tc.input.id)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestInMemoryAdRepository_Slice(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		name     string
		input    map[string]domain.Ad
		expected []domain.Ad
	}{
		{
			"All Ads",
			map[string]domain.Ad{
				"1": {
					Id:          "1",
					Title:       "Title 1",
					Description: "Description 1",
					Price:       45,
					Date:        now,
				},
				"2": {
					Id:          "2",
					Title:       "Title 2",
					Description: "Description 2",
					Price:       19,
					Date:        now,
				},
			},
			[]domain.Ad{
				{
					Id:          "1",
					Title:       "Title 1",
					Description: "Description 1",
					Price:       45,
					Date:        now,
				},
				{
					Id:          "2",
					Title:       "Title 2",
					Description: "Description 2",
					Price:       19,
					Date:        now,
				},
			},
		},
		{
			"No Ads",
			map[string]domain.Ad{},
			[]domain.Ad{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var inMemoryAdRepository inMemoryAdRepository = tc.input

			actual := inMemoryAdRepository.Slice()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func compareAds(actualValue domain.Ad, expectedValue domain.Ad) func() bool {
	return func() bool {
		return actualValue.Id == expectedValue.Id &&
			actualValue.Title == expectedValue.Title &&
			actualValue.Price == expectedValue.Price &&
			actualValue.Date != time.Time{}
	}
}
