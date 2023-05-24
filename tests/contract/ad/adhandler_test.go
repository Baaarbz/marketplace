package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server"
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"barbz.dev/marketplace/internal/pkg/application/ad/mocks"
	domainMock "barbz.dev/marketplace/internal/pkg/domain/ad/mocks"
	"context"
	"os"
	"testing"
	"time"
)

var (
	srv            server.Server
	findAllAdsMock = new(mocks.FindAllAds)
	findAdByIdMock = new(mocks.FindAdById)
	saveAdMock     = new(mocks.SaveAd)
)

func TestMain(m *testing.M) {
	adConfig := &configuration.AdConfiguration{
		Ads:               new(domainMock.AdRepository),
		SaveAdService:     saveAdMock,
		FindAllAdsService: findAllAdsMock,
		FindAdByIdService: findAdByIdMock,
	}
	_, srv = server.New(context.Background(), "localhost", 8080, 10*time.Second, adConfig)
	exitCode := m.Run()
	os.Exit(exitCode)
}
