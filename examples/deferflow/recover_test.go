package deferflow

import "testing"

func TestRecoverはDeferした関数から直接呼ばなければPanicを回復できない(t *testing.T) {
	directlyRecovered, helperRecovered, panicEscaped := RecoverResults()

	if !directlyRecovered {
		t.Fatal("directlyRecovered = false, want true")
	}
	if helperRecovered {
		t.Fatal("helperRecovered = true, want false")
	}
	if !panicEscaped {
		t.Fatal("panicEscaped = false, want true")
	}
}
