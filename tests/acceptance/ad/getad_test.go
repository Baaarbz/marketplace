package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server/handler/ad"
	"barbz.dev/marketplace/tests/acceptance"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindAllAds(t *testing.T) {
	var response []ad.JSONFindAllAdsResponse
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ads", nil)
	acceptance.Srv.Engine.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 200, w.Code)
	assert.True(t, len(response) > 1)
	assert.Contains(t, response, ad.JSONFindAllAdsResponse{Id: "90bc793b-3419-4b0a-9101-7ff4e7c12664"})
	assert.Contains(t, response, ad.JSONFindAllAdsResponse{Id: "ef483c32-ca95-47c9-8643-c3f23706ee4c"})
	assert.Contains(t, response, ad.JSONFindAllAdsResponse{Id: "c5fc62e5-5eea-40dd-a532-6262e0bec55a"})
	assert.Contains(t, response, ad.JSONFindAllAdsResponse{Id: "87dd12af-051e-4567-ac87-43361df0bf81"})
	assert.Contains(t, response, ad.JSONFindAllAdsResponse{Id: "2502bfa6-af82-427c-a8c6-e73d84f4d7a7"})
}

func TestFindAdById(t *testing.T) {
	var response ad.JSONFindAdByIdResponse
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ads/2502bfa6-af82-427c-a8c6-e73d84f4d7a7", nil)
	acceptance.Srv.Engine.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "2502bfa6-af82-427c-a8c6-e73d84f4d7a7", response.Id)
	assert.Equal(t, "Northface t-shirt", response.Title)
	assert.Equal(t, "This is the description to test the app", response.Description)
	assert.Equal(t, float32(799.50), response.Price)
	assert.NotNil(t, response.Date)
}
