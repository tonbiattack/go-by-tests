package stringexample

import "testing"

func Test書き込み後のstringsBuilderをコピーするとpanicになる(t *testing.T) {
	if recovered := CopyAfterWritePanics(); recovered == nil {
		t.Fatal("CopyAfterWritePanics() did not panic")
	}
}
