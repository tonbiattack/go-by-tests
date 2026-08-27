package timevalue

import "testing"

func Test同じ時刻でもlocationが異なるtimeTimeはEqualと比較結果が異なる(t *testing.T) {
	directEqual, instantEqual := CompareSameInstant()
	if directEqual || !instantEqual {
		t.Fatalf("directEqual = %t, instantEqual = %t", directEqual, instantEqual)
	}
}
