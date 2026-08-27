package errorexample

import "testing"

func TestラップしたerrorはerrorsIsで原因と照合できる(t *testing.T) {
	err := WrappedNotFound()

	if err == ErrNotFound {
		t.Fatal("wrapped error should not be directly equal to its cause")
	}
	if !IsNotFound(err) {
		t.Fatal("errors.Is should find the wrapped cause")
	}
}
