package methods

import "testing"

func TestValueReceiverはコピーを更新しpointerReceiverは呼び出し元を更新する(t *testing.T) {
	counter := Counter{}
	counter.IncrementValue()
	if counter.Value != 0 {
		t.Fatalf("value receiver changed counter: %d", counter.Value)
	}
	counter.IncrementPointer()
	if counter.Value != 1 {
		t.Fatalf("pointer receiver did not change counter: %d", counter.Value)
	}
}
