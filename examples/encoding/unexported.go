package encodingexample

import "encoding/json"

type Account struct {
	Name   string `json:"name"`
	secret string `json:"secret"`
}

func MarshalAccount() (string, error) {
	encoded, err := json.Marshal(Account{Name: "gopher", secret: "hidden"})
	return string(encoded), err
}
