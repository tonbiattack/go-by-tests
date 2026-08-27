package errorexample

import "testing"

func TestErrorsAsでポインタ型の原因を取り出せる(t *testing.T) {
	if got := FieldName(WrappedFieldError()); got != "email" {
		t.Fatalf("FieldName() = %q, want email", got)
	}
}
