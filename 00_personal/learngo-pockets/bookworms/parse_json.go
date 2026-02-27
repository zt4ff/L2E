package main

import (
	"encoding/json"
	"os"
)

func parseJSON[T any](path string) (T, error) {
	var result T

	f, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
