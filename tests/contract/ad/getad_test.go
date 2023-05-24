package ad

import (
	"barbz.dev/marketplace/internal/pkg/application/ad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindAllAds(t *testing.T) {
	serviceExpectedResponse := []ad.GetAdsDto{
		{Id: "test-id"},
	}
	expectedAllAdsJsonResponse := `[{"id": "test-id"}]`
	findAllAdsMock.EXPECT().Execute(mock.Anything).Return(serviceExpectedResponse, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ads", nil)
	srv.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, expectedAllAdsJsonResponse, w.Body.String())
}

func TestFindAdById(t *testing.T) {
	serviceExpectedResponse := ad.GetAdByIdDto{
		Id:          "test-id",
		Title:       "iPhone",
		Description: "Description of the iPhone ad",
		Price:       20,
		Date:        "2200-12-01",
	}
	expectedAdByIdJsonResponse := `{"id": "test-id", "title": "iPhone", "description": "Description of the iPhone ad", "price": 20,  "postedAt":"2200-12-01"}`
	findAdByIdMock.EXPECT().Execute(mock.Anything, "test-id").Return(serviceExpectedResponse, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ads/test-id", nil)
	srv.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, expectedAdByIdJsonResponse, w.Body.String())
}
