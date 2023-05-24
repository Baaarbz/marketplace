package factory

import (
	"barbz.dev/marketplace/internal/infrastructure/server/handler/ad"
	domain "barbz.dev/marketplace/internal/pkg/domain/ad"
	"barbz.dev/marketplace/pkg/valueobject"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
)

func AJSONSaveAdRequest() (objectRequest ad.JSONSaveAdRequest, jsonRequest []byte) {
	car := gofakeit.Vehicle()

	objectRequest = ad.JSONSaveAdRequest{
		Title:       fmt.Sprintf("%s %s", car.Brand, car.Model),
		Description: "Fake description for the ad of the car",
		Price:       float32(gofakeit.Price(1000, 100000)),
	}
	jsonRequest, _ = json.Marshal(objectRequest)
	return
}

func AnAd() domain.Ad {
	car := gofakeit.Vehicle()
	title, _ := valueobject.NewTitle(fmt.Sprintf("%s %s", car.Brand, car.Model))
	description, _ := valueobject.NewDescription("Fake description for the ad of the car")
	price, _ := valueobject.NewPrice(float32(gofakeit.Price(1000, 100000)))

	return domain.NewAd(title, description, price)
}
