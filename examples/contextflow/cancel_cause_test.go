package contextflow

import (
	"context"
	"errors"
	"testing"
)

func TestWithCancelCauseではErrはCanceledでCauseは指定した原因になる(t *testing.T) {
	state, cause := CanceledWithCause()

	if !errors.Is(state, context.Canceled) {
		t.Fatalf("state = %v, want %v", state, context.Canceled)
	}
	if !errors.Is(cause, ErrUpstreamUnavailable) {
		t.Fatalf("cause = %v, want %v", cause, ErrUpstreamUnavailable)
	}
}
