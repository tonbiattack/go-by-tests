package contextflow

import (
	"context"
	"errors"
	"testing"
)

func TestTimeout期限切れはContextDeadlineExceededを返す(t *testing.T) {
	ctx, cancel := ExpiredContext()
	defer cancel()
	if err := WaitForDeadline(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("WaitForDeadline() = %v, want DeadlineExceeded", err)
	}
}
