package valueobject

import (
	"time"
)

type Date string

func NewDate(value time.Time) Date {
	return Date(value.Format(time.RFC3339))
}

func (date Date) Time() (time.Time, error) {
	return time.Parse(time.RFC3339, string(date))
}

func (date Date) String() string {
	return string(date)
}
