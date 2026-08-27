package mapexample

import "testing"

func TestNilMapは読み取りできても書き込みでpanicになる(t *testing.T) {
	if got := ReadNilMap(); got != 0 {
		t.Fatalf("ReadNilMap() = %d, want 0", got)
	}
	if recovered := WriteNilMapPanics(); recovered == nil {
		t.Fatal("WriteNilMapPanics() did not panic")
	}
}
