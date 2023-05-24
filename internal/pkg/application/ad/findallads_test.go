package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"barbz.dev/marketplace/internal/pkg/domain/ad/mocks"
	"barbz.dev/marketplace/pkg/valueobject"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFindAllAds_Execute_GetListOfAds(t *testing.T) {
	ads := mocks.NewAdRepository(t)
	service := NewFindAllAds(ads)

	anId, _ := valueobject.NewId("574cc928-f4bd-11ed-ad0e-8a6a68a798d6")
	ad := NewAd("Simple title", "Simple ad description for testing", 20)
	ad.SetId(anId)
	ads.EXPECT().FindAllAds(mock.AnythingOfType("*context.emptyCtx")).Return([]Ad{ad}, nil)

	gotAds, err := service.Execute(context.Background())

	assert.True(t, len(gotAds) == 1)
	assert.Nil(t, err)
	assert.Equal(t, []GetAdsDto{{anId.String()}}, gotAds)
}
