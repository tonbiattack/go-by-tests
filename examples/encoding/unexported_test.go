package encodingexample

import "testing"

func TestEncodingJsonは非公開フィールドを出力しない(t *testing.T) {
	encoded, err := MarshalAccount()
	if err != nil {
		t.Fatal(err)
	}
	if encoded != `{"name":"gopher"}` {
		t.Fatalf("MarshalAccount() = %s", encoded)
	}
}
