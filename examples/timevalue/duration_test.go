package timevalue

import (
	"testing"
	"time"
)

func Test無単位のDurationはナノ秒として扱われる(t *testing.T) {
	if got := TimeoutFromNumber(100); got != 100*time.Nanosecond {
		t.Fatalf("TimeoutFromNumber(100) = %s, want 100ns", got)
	}
	if got := TimeoutInMilliseconds(100); got != 100*time.Millisecond {
		t.Fatalf("TimeoutInMilliseconds(100) = %s, want 100ms", got)
	}
}
