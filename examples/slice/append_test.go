package sliceexample

import "testing"

func Test容量に余裕があるsliceへのappendは同じバッキング配列を書き換える(t *testing.T) {
	values := make([]string, 2, 3)
	values[0], values[1] = "first", "second"

	extended := AppendLabel(values)

	if got := values[:cap(values)][2]; got != "third" {
		t.Fatalf("backing array value = %q, want %q", got, "third")
	}
	if &values[0] != &extended[0] {
		t.Fatal("append allocated unexpectedly")
	}
}
