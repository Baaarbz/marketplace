package ad

import (
	adHandler "barbz.dev/marketplace/internal/infrastructure/server/handler/ad"
	"barbz.dev/marketplace/internal/pkg/application/ad"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveAd(t *testing.T) {
	serviceExpectedResponse := ad.SaveAdDtoResponse{Id: "test-id"}
	expectedSaveAdJsonResponse := `{"id": "test-id"}`
	body, _ := json.Marshal(adHandler.JSONSaveAdRequest{
		Title:       "iPhone",
		Description: "Description og the iPhone ad test",
		Price:       100,
	})
	saveAdMock.EXPECT().Execute(mock.Anything, mock.AnythingOfType("ad.SaveAdDtoRequest")).Return(serviceExpectedResponse, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/ads", bytes.NewReader(body))
	srv.Engine.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.JSONEq(t, expectedSaveAdJsonResponse, w.Body.String())
}
