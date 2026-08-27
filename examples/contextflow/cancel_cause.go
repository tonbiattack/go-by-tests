package contextflow

import (
	"context"
	"errors"
)

var ErrUpstreamUnavailable = errors.New("upstream unavailable")

func WaitForCancelCause(ctx context.Context) (state error, cause error) {
	<-ctx.Done()
	return ctx.Err(), context.Cause(ctx)
}

func CanceledWithCause() (state error, cause error) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(ErrUpstreamUnavailable)
	return WaitForCancelCause(ctx)
}
