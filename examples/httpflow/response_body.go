package httpflow

import (
	"io"
	"net/http"
)

func ReadAndClose(response *http.Response) (string, error) {
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	return string(bytes), err
}
