package ad

import (
	. "barbz.dev/marketplace/pkg/valueobject"
	"time"
)

type Ad struct {
	id          AdId
	Title       Title
	Description Description
	Price       Price
	date        Date
}

func NewAd(title Title, description Description, price Price) Ad {
	return Ad{
		Title:       title,
		Description: description,
		Price:       price,
		date:        NewDate(time.Now()),
	}
}

func (ad *Ad) GetId() AdId {
	return ad.id
}

func (ad *Ad) SetId(id AdId) {
	ad.id = id
}

func (ad *Ad) GetDate() Date {
	return ad.date
}
