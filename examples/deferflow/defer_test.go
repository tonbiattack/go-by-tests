package deferflow

import (
	"reflect"
	"testing"
)

func TestDeferの引数はdefer文を実行した時点で評価される(t *testing.T) {
	if got, want := CapturedArgument(), []string{"before"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("CapturedArgument() = %v, want %v", got, want)
	}
}

func TestDeferは後に登録した関数からLIFO順に実行される(t *testing.T) {
	if got, want := LIFO(), []int{2, 1}; !reflect.DeepEqual(got, want) {
		t.Fatalf("LIFO() = %v, want %v", got, want)
	}
}
