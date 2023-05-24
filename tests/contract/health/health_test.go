package health

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"barbz.dev/marketplace/internal/pkg/application/ad/mocks"
	domainMock "barbz.dev/marketplace/internal/pkg/domain/ad/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var srv server.Server

func TestMain(m *testing.M) {
	adConfig := &configuration.AdConfiguration{
		Ads:               new(domainMock.AdRepository),
		SaveAdService:     new(mocks.SaveAd),
		FindAllAdsService: new(mocks.FindAllAds),
		FindAdByIdService: new(mocks.FindAdById),
	}
	_, srv = server.New(context.Background(), "localhost", 8080, 10*time.Second, adConfig)
	exitCode := m.Run()
	os.Exit(exitCode)
}
func TestFindAdById(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	srv.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "everything is ok!", w.Body.String())
}
