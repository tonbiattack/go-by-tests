package mapexample

import "testing"

func Test存在しないmapキーも値だけを読むとゼロ値になる(t *testing.T) {
	scores := map[string]int{"zero": 0}

	missingScore, missingOK := Lookup(scores, "missing")
	zeroScore, zeroOK := Lookup(scores, "zero")
	if missingScore != 0 || missingOK || zeroScore != 0 || !zeroOK {
		t.Fatalf("missing = (%d, %t), zero = (%d, %t)", missingScore, missingOK, zeroScore, zeroOK)
	}
}
