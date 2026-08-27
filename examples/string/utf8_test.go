package stringexample

import "testing"

func Test絵文字を含む文字列ではバイト数とrune数が異なる(t *testing.T) {
	value := "Go🐹"

	if got := ByteLen(value); got != 6 {
		t.Fatalf("ByteLen() = %d, want 6", got)
	}
	if got := RuneLen(value); got != 3 {
		t.Fatalf("RuneLen() = %d, want 3", got)
	}
}
