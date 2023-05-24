package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
	"github.com/google/uuid"
)

type InMemoryRepository struct {
	ads []Ad
}

func NewInMemoryRepository() AdRepository {
	return &InMemoryRepository{
		ads: make([]Ad, 0),
	}
}

func (repository *InMemoryRepository) SaveAd(_ context.Context, ad Ad) (Ad, error) {
	var id, _ = uuid.NewUUID()
	var adId, _ = NewId(id.String())
	ad.SetId(adId)

	repository.ads = append(repository.ads, ad)
	return ad, nil
}

func (repository *InMemoryRepository) FindAdById(_ context.Context, id AdId) (Ad, error) {
	for _, ad := range repository.ads {
		if ad.GetId() == id {
			return ad, nil
		}
	}
	return Ad{}, nil
}

func (repository *InMemoryRepository) FindAllAds(_ context.Context) (adResponse []Ad, err error) {
	if len(repository.ads) < 5 {
		adResponse = repository.ads
	} else {
		adResponse = repository.ads[:5]
	}
	return
}
