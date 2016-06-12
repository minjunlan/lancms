package lan

import (
	"errors"
	"time"
)

type Error struct {
	When time.Time
	What string
}

func New(text string) error {
	return errors.New(text)
}
