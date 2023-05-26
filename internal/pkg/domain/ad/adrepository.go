//go:generate mockery --name=AdRepository --filename mock_adrepository.go
package ad

import (
	"barbz.dev/marketplace/pkg/valueobject"
	"context"
)

type AdRepository interface {
	SaveAd(ctx context.Context, ad Ad) (Ad, error)
	SaveAdPostedSite(ctx context.Context, adId valueobject.AdId, site string) error
	FindAdById(ctx context.Context, id valueobject.AdId) (Ad, error)
	FindAllAds(ctx context.Context) (adResponse []Ad, err error)
}
