package encodingexample

import "encoding/json"

type Summary struct {
	Name  string `json:"name"`
	Count int    `json:"count,omitempty"`
}

func Marshal(summary Summary) (string, error) {
	encoded, err := json.Marshal(summary)
	return string(encoded), err
}
