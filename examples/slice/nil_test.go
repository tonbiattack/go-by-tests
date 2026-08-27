package sliceexample

import "testing"

func TestNilSliceと空sliceはJSONで異なる値になる(t *testing.T) {
	nilJSON, err := Marshal(nil)
	if err != nil {
		t.Fatal(err)
	}
	emptyJSON, err := Marshal([]string{})
	if err != nil {
		t.Fatal(err)
	}

	if nilJSON != "null" || emptyJSON != "[]" {
		t.Fatalf("nil = %s, empty = %s", nilJSON, emptyJSON)
	}
}

func TestNilSliceだけがnilと比較してtrueになる(t *testing.T) {
	if !IsNil(nil) || IsNil([]string{}) {
		t.Fatal("nil slice and empty slice must be distinguished")
	}
}
