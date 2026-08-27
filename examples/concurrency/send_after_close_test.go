package concurrencyexample

import "testing"

func Test閉じたChannelへの送信はPanicになる(t *testing.T) {
	if got := SendAfterClosePanics(); got == nil {
		t.Fatal("SendAfterClosePanics() = nil, want recovered panic")
	}
}
