package concurrencyexample

import "testing"

func TestSyncOnceはPanicしても再実行しない(t *testing.T) {
	if got := CallsAfterOncePanic(); got != 1 {
		t.Fatalf("CallsAfterOncePanic() = %d, want 1", got)
	}
}
