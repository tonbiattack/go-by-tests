package stringexample

import (
	"reflect"
	"testing"
)

func Test文字列に対するrangeのindexはrune番号ではなくバイト位置になる(t *testing.T) {
	if got, want := RuneByteIndexes("A🐹B"), []int{0, 1, 5}; !reflect.DeepEqual(got, want) {
		t.Fatalf("RuneByteIndexes() = %v, want %v", got, want)
	}
}
