package errorexample

import (
	"errors"
	"testing"
)

func TestErrorsJoinした複数の原因をErrorsIsで照合できる(t *testing.T) {
	err := CombinedFailure()
	if !errors.Is(err, ErrDatabase) || !errors.Is(err, ErrTimeout) {
		t.Fatalf("errors.Is did not find every joined error: %v", err)
	}
}
