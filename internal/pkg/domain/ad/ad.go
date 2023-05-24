package ad

import (
	. "barbz.dev/marketplace/pkg/valueobject"
	"time"
)

type Ad struct {
	Id          AdId
	Title       Title
	Description Description
	Price       Price
	Date        time.Time
}

func NewAd(title Title, description Description, price Price, date time.Time) Ad {
	return Ad{
		Title:       title,
		Description: description,
		Price:       price,
		Date:        date,
	}
}
