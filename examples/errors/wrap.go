package errorexample

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func WrappedNotFound() error {
	return fmt.Errorf("load profile: %w", ErrNotFound)
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
