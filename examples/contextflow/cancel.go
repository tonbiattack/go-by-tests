package contextflow

import "context"

func WaitForCancel(ctx context.Context) error {
	<-ctx.Done()
	return ctx.Err()
}
