//go:generate mockery --name=FindAdById --filename mock_findadbyid.go
package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"barbz.dev/marketplace/pkg/valueobject"
	"context"
)

type FindAdById interface {
	Execute(ctx context.Context, id string) (response GetAdByIdDto, err error)
}
type findAdById struct {
	ads AdRepository
}

func NewFindAdById(ads AdRepository) FindAdById {
	return findAdById{
		ads: ads,
	}
}

type GetAdByIdDto struct {
	Id          string
	Title       string
	Description string
	Price       float32
	Date        string
}

func (service findAdById) Execute(ctx context.Context, id string) (response GetAdByIdDto, err error) {
	adId, err := valueobject.NewId(id)
	if err != nil {
		return
	}
	ad, err := service.ads.FindAdById(ctx, adId)
	response = service.mapToResponse(ad)
	return
}

func (findAdById) mapToResponse(ad Ad) GetAdByIdDto {
	return GetAdByIdDto{
		Id:          ad.GetId().String(),
		Title:       ad.Title.String(),
		Description: ad.Description.String(),
		Price:       ad.Price.Number(),
		Date:        ad.GetDate().String(),
	}
}
