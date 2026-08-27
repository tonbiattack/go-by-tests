package timevalue

import "time"

func TimeoutFromNumber(value time.Duration) time.Duration {
	return value
}

func TimeoutInMilliseconds(value int) time.Duration {
	return time.Duration(value) * time.Millisecond
}
