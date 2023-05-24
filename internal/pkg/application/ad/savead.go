//go:generate mockery --name=SaveAd --filename mock_savead.go
package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
)

type SaveAd interface {
	Execute(ctx context.Context, request SaveAdDtoRequest) (SaveAdDtoResponse, error)
}

type saveAd struct {
	ads AdRepository
}

func NewSaveAd(ads AdRepository) SaveAd {
	return saveAd{
		ads: ads,
	}
}

type SaveAdDtoResponse struct {
	Id string
}

type SaveAdDtoRequest struct {
	Title       string
	Description string
	Price       float32
}

func (service saveAd) Execute(ctx context.Context, request SaveAdDtoRequest) (SaveAdDtoResponse, error) {
	if title, description, price, err := service.getFieldsAds(request); err != nil {
		return SaveAdDtoResponse{}, err
	} else {
		ad := NewAd(title, description, price)
		ad, err = service.ads.SaveAd(ctx, ad)
		return SaveAdDtoResponse{Id: ad.GetId().String()}, err
	}
}

func (saveAd) getFieldsAds(request SaveAdDtoRequest) (title Title, description Description, price Price, err error) {
	title, err = NewTitle(request.Title)
	if err != nil {
		return "", "", 0, err
	}
	description, err = NewDescription(request.Description)
	if err != nil {
		return "", "", 0, err
	}
	price, err = NewPrice(request.Price)
	return
}
