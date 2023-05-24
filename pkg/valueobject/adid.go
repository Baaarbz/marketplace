package valueobject

import (
	"errors"
	"github.com/google/uuid"
)

var ErrAdIdBadFormat = errors.New("id: bad format")

type AdId string

func NewId(value string) (adId AdId, err error) {
	if err = validateAdId(value); err != nil {
		return "", err
	}
	adId = AdId(value)
	return
}

func (id AdId) String() string {
	return string(id)
}

func validateAdId(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return ErrAdIdBadFormat
	}

	return nil
}
