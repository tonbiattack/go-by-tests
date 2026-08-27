package encodingexample

import (
	"encoding/json"
	"strings"
)

type CreateUserInput struct {
	Name string `json:"name"`
}

func DecodeIgnoringUnknown(input string) (CreateUserInput, error) {
	var user CreateUserInput
	err := json.Unmarshal([]byte(input), &user)
	return user, err
}

func DecodeRejectingUnknown(input string) (CreateUserInput, error) {
	var user CreateUserInput
	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	return user, err
}
