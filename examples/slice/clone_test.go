package sliceexample

import "testing"

func TestSlicesCloneは別のバッキング配列を作る(t *testing.T) {
	original, cloned := CloneThenChange([]int{1, 2})
	if original[0] != 1 || cloned[0] != 99 {
		t.Fatalf("original = %v, cloned = %v", original, cloned)
	}
}
