package concurrencyexample

import "testing"

func TestBufferedChannelは受信者なしでも容量まで送信できる(t *testing.T) {
	if SendWithoutReceiver(0) {
		t.Fatal("unbuffered channel sent without receiver")
	}
	if !SendWithoutReceiver(1) {
		t.Fatal("buffered channel did not accept value")
	}
}
