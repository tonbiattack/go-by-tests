package httpflow

import (
	"net/http"
	"strings"
	"testing"
)

type trackedBody struct {
	*strings.Reader
	closed bool
}

func (body *trackedBody) Close() error {
	body.closed = true
	return nil
}

func TestHTTPResponseBodyは読み終えた後にCloseする(t *testing.T) {
	body := &trackedBody{Reader: strings.NewReader("go")}
	response := &http.Response{Body: body}
	content, err := ReadAndClose(response)
	if err != nil {
		t.Fatal(err)
	}
	if content != "go" || !body.closed {
		t.Fatalf("content = %q, closed = %t", content, body.closed)
	}
}
