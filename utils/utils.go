package utils

import (
	"encoding/json"
	"net/http"
)

type Utils interface {
	Fetch(any) error
}

func NewUtils() utils {
	return utils{}
}

type utils struct {
}

func (ut utils) Fetch(result any) error {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return err
	}

	defer response.Body.Close()

	decodeErr := json.NewDecoder(response.Body).Decode(&result)
	if decodeErr != nil {
		return decodeErr
	}
	return nil
}
