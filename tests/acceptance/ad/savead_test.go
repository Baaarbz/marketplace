package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server/handler/ad"
	"barbz.dev/marketplace/pkg/valueobject"
	"barbz.dev/marketplace/tests/factory"
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveAd(t *testing.T) {
	objectRequest, jsonRequest := factory.AJSONSaveAdRequest()
	var response ad.JSONSaveAdResponse
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/ads", bytes.NewReader(jsonRequest))
	srv.Engine.ServeHTTP(w, req)

	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 201, w.Code)
	result, _ := dependencies.Ads.FindAdById(context.Background(), valueobject.AdId(response.Id))
	assert.Equal(t, objectRequest.Description, result.Description.String())
	assert.Equal(t, objectRequest.Title, result.Title.String())
	assert.Equal(t, objectRequest.Price, result.Price.Number())
	assert.NotNil(t, result.GetDate())
	assert.NotNil(t, result.GetId())
}
