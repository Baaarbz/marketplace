package configuration

import (
	repository "barbz.dev/marketplace/internal/infrastructure/repository/ad"
	service "barbz.dev/marketplace/internal/pkg/application/ad"
	domain "barbz.dev/marketplace/internal/pkg/domain/ad"
	"go.uber.org/dig"
)

type AdConfiguration struct {
	Ads               domain.AdRepository
	SaveAdService     service.SaveAd
	FindAllAdsService service.FindAllAds
	FindAdByIdService service.FindAdById
}

func BuildAdConfiguration() (*AdConfiguration, error) {
	container, err := buildContainer()
	if err != nil {
		return nil, err
	}

	dependencies := &AdConfiguration{}
	if err := container.Invoke(func(
		ads domain.AdRepository,
		saveAdService service.SaveAd,
		findAllAdsService service.FindAllAds,
		findAdByIdService service.FindAdById,
	) {
		dependencies.Ads = ads
		dependencies.SaveAdService = saveAdService
		dependencies.FindAdByIdService = findAdByIdService
		dependencies.FindAllAdsService = findAllAdsService
	}); err != nil {
		return nil, err
	}

	return dependencies, nil
}

func buildContainer() (*dig.Container, error) {
	container := dig.New()
	if err := container.Provide(repository.NewInMemoryRepository); err != nil {
		return nil, err
	}

	if err := container.Provide(service.NewSaveAd); err != nil {
		return nil, err
	}

	if err := container.Provide(service.NewFindAllAds); err != nil {
		return nil, err
	}

	if err := container.Provide(service.NewFindAdById); err != nil {
		return nil, err
	}

	return container, nil
}
