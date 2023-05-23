package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/javier-porto/learning-go/application"
	"github.mpi-internal.com/javier-porto/learning-go/domain"
	"github.mpi-internal.com/javier-porto/learning-go/infrastructure/client"
	"github.mpi-internal.com/javier-porto/learning-go/infrastructure/repository"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPostAd(t *testing.T) {
	adRepository := repository.NewInMemoryAdRepository()
	server := client.SetupServer(
		application.NewAdService(
			adRepository,
		),
	)

	t.Run("API persists the ad entity from the request", func(t *testing.T) {
		httpRequest := client.HttpCreateAdRequest{
			Title:       "Acceptance Test Ad",
			Description: "Ad for acceptance test purposes",
			Price:       59,
		}
		jsonValue, _ := json.Marshal(httpRequest)
		req, _ := http.NewRequest("POST", "/ads", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("API returns Bad Request if request does not comply with the contract", func(t *testing.T) {
		jsonValue := []byte(`{"Description": "Desc", "Price": 59, "Foo": "Bar"}`)
		req, _ := http.NewRequest("POST", "/ads", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("API returns Conflict if an ad with the same title already exists", func(t *testing.T) {
		adRepository.Persist(domain.Ad{
			Id:          "cfec58d9-8cf7-40a3-b0f9-a6d1b80df273",
			Title:       "Already Exists",
			Description: "Desc",
			Price:       99,
			Date:        time.Now(),
		})

		httpRequest := client.HttpCreateAdRequest{
			Title:       "Already Exists",
			Description: "Ad for acceptance test purposes",
			Price:       37,
		}
		jsonValue, _ := json.Marshal(httpRequest)
		req, _ := http.NewRequest("POST", "/ads", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		var responseBody client.HttpErrorResponse
		json.Unmarshal(w.Body.Bytes(), &responseBody)
		expectedResponse := client.HttpErrorResponse{
			Code:  409,
			Title: "ad-already-exists",
			Error: fmt.Sprintf("an ad with title %s already exists", httpRequest.Title),
		}
		assert.Equal(t, expectedResponse, responseBody)
	})
}

func TestGetAd(t *testing.T) {
	adRepository := repository.NewInMemoryAdRepository()
	now := time.Now()
	adRepository.Persist(domain.Ad{
		Id:          "3c8a884b-9c80-45c4-bb85-7945182034fb",
		Title:       "Acceptance Test Ad",
		Description: "Ad for acceptance test purposes",
		Price:       59,
		Date:        now,
	})
	server := client.SetupServer(
		application.NewAdService(
			adRepository,
		),
	)

	t.Run("API returns the requested ad", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/ads/3c8a884b-9c80-45c4-bb85-7945182034fb", nil)

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		responseStatus := w.Code
		assert.Equal(t, http.StatusOK, responseStatus)
		var responseBody client.HttpAdResponse
		json.Unmarshal(w.Body.Bytes(), &responseBody)
		expectedResponse := client.HttpAdResponse{
			Id:          "3c8a884b-9c80-45c4-bb85-7945182034fb",
			Title:       "Acceptance Test Ad",
			Description: "Ad for acceptance test purposes",
			Price:       59,
			Date:        now,
		}
		assert.Equal(t, expectedResponse.Id, responseBody.Id)
		assert.Equal(t, expectedResponse.Title, responseBody.Title)
		assert.Equal(t, expectedResponse.Description, responseBody.Description)
		assert.Equal(t, expectedResponse.Price, responseBody.Price)
	})
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}
