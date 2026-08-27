package interfacevalue

import (
	"bytes"
	"testing"
)

func TestNilポインタを入れたinterfaceはnilではない(t *testing.T) {
	var buffer *bytes.Buffer

	if IsNilReader(buffer) {
		t.Fatal("typed nil stored in an interface must not compare equal to nil")
	}
}
