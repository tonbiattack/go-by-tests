package sliceexample

import "encoding/json"

func Marshal(values []string) (string, error) {
	encoded, err := json.Marshal(values)
	return string(encoded), err
}

func IsNil(values []string) bool {
	return values == nil
}
