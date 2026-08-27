package encodingexample

import "testing"

func TestEncodingJsonは未知フィールドを既定で無視しDisallowUnknownFieldsで拒否できる(t *testing.T) {
	const input = `{"name":"gopher","role":"admin"}`

	ignored, err := DecodeIgnoringUnknown(input)
	if err != nil {
		t.Fatalf("DecodeIgnoringUnknown() error = %v", err)
	}
	if ignored.Name != "gopher" {
		t.Fatalf("DecodeIgnoringUnknown().Name = %q, want %q", ignored.Name, "gopher")
	}

	if _, err := DecodeRejectingUnknown(input); err == nil {
		t.Fatal("DecodeRejectingUnknown() error = nil, want unknown-field error")
	}
}
