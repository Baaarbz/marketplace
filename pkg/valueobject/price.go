package valueobject

import (
	"errors"
	"fmt"
)

var ErrPriceBadFormat = errors.New("price: bad format")

type Price float32

func NewPrice(value float32) (adPrice Price, err error) {
	if err = validateAdPrice(value); err != nil {
		return 0, err
	}
	adPrice = Price(value)
	return
}

func (price Price) Number() float32 {
	return float32(price)
}

func (price Price) String() string {
	return fmt.Sprint(float32(price))
}

func validateAdPrice(price float32) error {
	if price < 0 {
		return ErrPriceBadFormat
	}

	return nil
}
