package valueobject

import (
	"errors"
	"strings"
)

var ErrDescriptionBadFormat = errors.New("description: bad format")

type Description string

func NewDescription(value string) (adDescription Description, err error) {
	if err = validateAdDescription(value); err != nil {
		return "", err
	}
	adDescription = Description(value)
	return
}

func (description Description) String() string {
	return string(description)
}

func validateAdDescription(description string) error {
	words := strings.Fields(description)
	if descriptionLen := len(words); descriptionLen <= 3 {
		return ErrDescriptionBadFormat
	}

	return nil
}
