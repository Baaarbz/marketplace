package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryRepository_FindAdById(t *testing.T) {
	anId, _ := NewId("574cc928-f4bd-11ed-ad0e-8a6a68a798d6")
	anAd := mockAdWithId(anId)

	type fields struct {
		ads []Ad
	}
	type args struct {
		id AdId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantAd  Ad
		wantErr bool
	}{
		{"empty slice should empty ad", fields{ads: []Ad{}}, args{id: anId}, Ad{}, false},
		{"not find ad by id should empty ad", fields{ads: generateSliceMockAds()}, args{id: anId}, Ad{}, false},
		{"slice with values should find ad", fields{ads: append(generateSliceMockAds(), anAd)}, args{id: anId}, anAd, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &InMemoryRepository{
				ads: tt.fields.ads,
			}
			gotAd, err := repository.FindAdById(context.Background(), tt.args.id)

			assert.Equal(t, tt.wantAd, gotAd)
			switch tt.wantErr {
			case true:
				assert.NotNil(t, err)
			case false:
				assert.Nil(t, err)
			}
		})
	}
}

func TestInMemoryRepository_FindAllAds(t *testing.T) {
	sliceAdsMock := generateSliceMockAds()
	type fields struct {
		ads []Ad
	}
	tests := []struct {
		name           string
		fields         fields
		wantAdResponse []Ad
	}{
		{"empty ads slice should return empty slice", fields{ads: []Ad{}}, []Ad{}},
		{"slice with 4 elements should return 4 elements", fields{ads: sliceAdsMock[:4]}, sliceAdsMock[:4]},
		{"slice with more than 5 elements should return 5 elements", fields{ads: sliceAdsMock}, sliceAdsMock[:5]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &InMemoryRepository{
				ads: tt.fields.ads,
			}
			gotAdResponse, err := repository.FindAllAds(context.Background())
			assert.Nil(t, err)
			assert.Equal(t, tt.wantAdResponse, gotAdResponse)
		})
	}
}

func TestInMemoryRepository_SaveAd(t *testing.T) {
	anAd := NewAd("Test Save", "Test save Ad description mock", 10)
	repository := &InMemoryRepository{ads: []Ad{}}

	expectedAd, err := repository.SaveAd(context.Background(), anAd)

	gotAd, _ := repository.FindAdById(context.Background(), expectedAd.GetId())
	assert.Equal(t, expectedAd, gotAd)
	assert.Nil(t, err)
}

func generateSliceMockAds() (mockAds []Ad) {
	mockAds = make([]Ad, 0)
	mockAds = append(mockAds, NewAd("Sample Ad", "This is the description to test the app", 30))
	mockAds = append(mockAds, NewAd("TV 45' Sony", "This is the description to test the app", 599.99))
	mockAds = append(mockAds, NewAd("Sportiva Rock Climbing Foot", "This is the description to test the app", 70))
	mockAds = append(mockAds, NewAd("Macbook Pro 16'", "This is the description to test the app", 1799.00))
	mockAds = append(mockAds, NewAd("Rolex limited edition", "This is the description to test the app", 100500))
	mockAds = append(mockAds, NewAd("Northface t-shirt", "This is the description to test the app", 55))

	return
}
func mockAdWithId(anId AdId) Ad {
	adWithId := NewAd("BMW 320D", "This is the description to test the app", 18500)
	adWithId.SetId(anId)

	return adWithId
}
