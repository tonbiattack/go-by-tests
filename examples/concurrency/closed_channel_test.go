package concurrencyexample

import "testing"

func Test閉じたchannelの受信はゼロ値とfalseを返す(t *testing.T) {
	first, afterClose, open := ReadAfterClose()
	if first != 7 || afterClose != 0 || open {
		t.Fatalf("got (%d, %d, %t), want (7, 0, false)", first, afterClose, open)
	}
}
