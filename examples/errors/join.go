package errorexample

import "errors"

var (
	ErrDatabase = errors.New("database unavailable")
	ErrTimeout  = errors.New("request timed out")
)

func CombinedFailure() error {
	return errors.Join(ErrDatabase, ErrTimeout)
}
