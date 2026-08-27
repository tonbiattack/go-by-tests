package deferflow

import (
	"errors"
	"testing"
)

func TestDeferしたCloseのエラーは明示的に返さなければ失われる(t *testing.T) {
	if err := IgnoreCloseError(); err != nil {
		t.Fatalf("IgnoreCloseError() = %v, want nil", err)
	}
	if err := ReturnCloseError(); !errors.Is(err, ErrClose) {
		t.Fatalf("ReturnCloseError() = %v, want ErrClose", err)
	}
}
