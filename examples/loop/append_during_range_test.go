package loopexample

import (
	"reflect"
	"testing"
)

func TestRange中にAppendしても反復回数は開始時の長さで固定される(t *testing.T) {
	visited, final := VisitWhileAppending()

	if want := []int{1, 2, 3}; !reflect.DeepEqual(visited, want) {
		t.Fatalf("visited = %v, want %v", visited, want)
	}
	if want := []int{1, 2, 3, 11, 12, 13}; !reflect.DeepEqual(final, want) {
		t.Fatalf("final = %v, want %v", final, want)
	}
}
