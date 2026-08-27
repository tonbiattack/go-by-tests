package errorexample

import (
	"errors"
	"fmt"
)

type FieldError struct {
	Field string
}

func (e *FieldError) Error() string {
	return e.Field + " is invalid"
}

func WrappedFieldError() error {
	return fmt.Errorf("request rejected: %w", &FieldError{Field: "email"})
}

func FieldName(err error) string {
	var target *FieldError
	if errors.As(err, &target) {
		return target.Field
	}
	return ""
}
