package contextflow

import (
	"context"
	"testing"
)

func TestCancelしたcontextはcontextCanceledを返す(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if got := WaitForCancel(ctx); got != context.Canceled {
		t.Fatalf("WaitForCancel() = %v, want %v", got, context.Canceled)
	}
}
