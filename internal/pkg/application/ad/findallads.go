//go:generate mockery --name=FindAllAds --filename mock_findallads.go
package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"context"
)

type FindAllAds interface {
	Execute(ctx context.Context) ([]GetAdsDto, error)
}

type findAllAds struct {
	ads AdRepository
}

func NewFindAllAds(ads AdRepository) FindAllAds {
	return findAllAds{
		ads: ads,
	}
}

type GetAdsDto struct {
	Id string
}

func (service findAllAds) Execute(ctx context.Context) ([]GetAdsDto, error) {
	ads, err := service.ads.FindAllAds(ctx)
	return service.mapToResponse(ads), err
}

func (findAllAds) mapToResponse(ads []Ad) []GetAdsDto {
	adsResponse := make([]GetAdsDto, 0)
	for _, ad := range ads {
		adsResponse = append(adsResponse, GetAdsDto{Id: ad.GetId().String()})
	}
	return adsResponse
}
