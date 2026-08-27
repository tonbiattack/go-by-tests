package encodingexample

import "testing"

func TestOmitEmptyは数値のゼロ値をJSONから省略する(t *testing.T) {
	emptyCount, err := Marshal(Summary{Name: "go"})
	if err != nil {
		t.Fatal(err)
	}
	countOne, err := Marshal(Summary{Name: "go", Count: 1})
	if err != nil {
		t.Fatal(err)
	}

	if emptyCount != `{"name":"go"}` || countOne != `{"name":"go","count":1}` {
		t.Fatalf("empty = %s, one = %s", emptyCount, countOne)
	}
}
