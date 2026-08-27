package contextflow

import (
	"context"
	"time"
)

func WaitForDeadline(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}

func ExpiredContext() (context.Context, context.CancelFunc) {
	return context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
}
