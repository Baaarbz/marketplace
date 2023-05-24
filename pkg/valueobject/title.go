package valueobject

import (
	"errors"
)

var ErrTitleBadFormat = errors.New("title: bad format")

type Title string

func NewTitle(value string) (adTitle Title, err error) {
	if err = validateAdTitle(value); err != nil {
		return "", err
	}
	adTitle = Title(value)
	return
}

func (title Title) String() string {
	return string(title)
}

func validateAdTitle(title string) error {
	if titleLen := len(title); titleLen <= 3 {
		return ErrTitleBadFormat
	}

	return nil
}
