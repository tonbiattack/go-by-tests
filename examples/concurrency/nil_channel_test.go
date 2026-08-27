package concurrencyexample

import "testing"

func TestNilChannelを含むSelectCaseは無効になる(t *testing.T) {
	if !NilChannelCaseIsDisabled() {
		t.Fatal("nil channel case must not be selected")
	}
}
