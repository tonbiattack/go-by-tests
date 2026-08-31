package values

import "testing"

func Test配列は値としてコピーされsliceは要素を共有する(t *testing.T) {
	array := [2]string{"original", "second"}
	slice := []string{"original", "second"}

	changedArray := ChangeArray(array)
	ChangeSlice(slice)

	if array[0] != "original" || changedArray[0] != "changed" || slice[0] != "changed" {
		t.Fatalf("array=%v changedArray=%v slice=%v", array, changedArray, slice)
	}
}
