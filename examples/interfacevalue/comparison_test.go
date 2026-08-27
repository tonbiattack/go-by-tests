package interfacevalue

import "testing"

func TestSliceを持つinterface値同士の比較はpanicになる(t *testing.T) {
	if recovered := CompareSlicesAsAnyPanics(); recovered == nil {
		t.Fatal("CompareSlicesAsAnyPanics() did not panic")
	}
}
