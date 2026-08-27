package stringexample

import "testing"

func TestStringをByteSliceへ変換して変更しても元の文字列は変わらない(t *testing.T) {
	unchanged, changed := ChangeFirstByte("go")
	if unchanged != "go" || changed != "Go" {
		t.Fatalf("original = %q, changed = %q", unchanged, changed)
	}
}
