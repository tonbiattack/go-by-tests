package concurrencyexample

import (
	"sync"
	"testing"
)

func TestSyncMapのLoad結果は型アサーションなしに使えない(t *testing.T) {
	var store sync.Map
	store.Store("count", 3)

	count, ok := LoadCount(&store, "count")
	if !ok || count != 3 {
		t.Fatalf("LoadCount() = %d, %t; want 3, true", count, ok)
	}

	store.Store("count", "three")
	if _, ok := LoadCount(&store, "count"); ok {
		t.Fatal("LoadCount() accepted a string value")
	}
}
