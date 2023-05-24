package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"barbz.dev/marketplace/internal/pkg/domain/ad/mocks"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFindAdById_Execute_FindAd(t *testing.T) {
	ads := mocks.NewAdRepository(t)
	service := NewFindAdById(ads)

	anId, _ := NewId("574cc928-f4bd-11ed-ad0e-8a6a68a798d6")
	ad := NewAd("Simple title", "Simple ad description for testing", 20)
	ad.SetId(anId)
	expectedResponse := GetAdByIdDto{
		Id:          anId.String(),
		Title:       ad.Title.String(),
		Description: ad.Description.String(),
		Price:       ad.Price.Number(),
		Date:        ad.GetDate().String(),
	}
	ads.EXPECT().FindAdById(mock.AnythingOfType("*context.emptyCtx"), anId).Return(ad, nil)

	gotAd, _ := service.Execute(context.Background(), anId.String())

	assert.Equal(t, expectedResponse, gotAd)
}

func TestFindAdById_Execute_NotFound(t *testing.T) {
	ads := mocks.NewAdRepository(t)
	service := NewFindAdById(ads)

	anId, _ := NewId("574cc928-f4bd-11ed-ad0e-8a6a68a798d6")
	ads.EXPECT().FindAdById(mock.AnythingOfType("*context.emptyCtx"), anId).Return(Ad{}, errors.New("test: error"))

	gotAd, err := service.Execute(context.Background(), anId.String())

	assert.Empty(t, gotAd)
	assert.NotNil(t, err)
}
