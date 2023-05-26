//go:generate mockery --name=SaveAd --filename mock_savead.go
package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
	"log"
	"math/rand"
	"time"
)

type SaveAd interface {
	Execute(ctx context.Context, request SaveAdDtoRequest) (SaveAdDtoResponse, error)
}

type saveAd struct {
	ads AdRepository
}

var sites = []string{"Milanuncios", "Subito", "Labencoin", "Kleinanzeigen", "Marktplaats", "OLX"}

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

func (s saveAd) Execute(ctx context.Context, request SaveAdDtoRequest) (SaveAdDtoResponse, error) {
	if title, description, price, err := s.getFieldsAds(request); err != nil {
		return SaveAdDtoResponse{}, err
	} else {
		ad := NewAd(title, description, price)
		ad, err = s.ads.SaveAd(ctx, ad)
		go s.findSitesWhereTheAdWasPosted(ctx, ad.Id)
		return SaveAdDtoResponse{Id: ad.Id.String()}, err
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

func (s saveAd) findSitesWhereTheAdWasPosted(ctx context.Context, adId AdId) {
	rand.NewSource(time.Now().UnixNano())

	amountOfSitesPosted := rand.Intn(3)
	for i := 0; i < amountOfSitesPosted; i++ {
		site := sites[rand.Intn(6)]
		if err := s.ads.SaveAdPostedSite(ctx, adId, site); err != nil {
			log.Fatalf("Something goes wrong trying to search where the ad was posted in other sites: %s", err)
		}
	}

}
