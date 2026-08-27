package loopexample

import (
	"reflect"
	"testing"
)

func TestGo122ではrange変数をgoroutineが各反復ごとに捕捉する(t *testing.T) {
	if got, want := RangeValuesFromGoroutines(), []int{0, 1, 2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("RangeValuesFromGoroutines() = %v, want %v", got, want)
	}
}
