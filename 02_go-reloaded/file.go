package main

import (
	"fmt"
	"os"
)

type wrongArgumentError string

func (w wrongArgumentError) Error() string {
	return string(wrongArgumentError(w))
}

const WrongArgumentError = wrongArgumentError("Wrong Argument Format")

func parseArgument() ([]string, error) {
	args := os.Args[1:]

	if len(args) != 2 {
		return nil, fmt.Errorf("%w: Argument provided must be 2", WrongArgumentError)
	}

	return args, nil
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func writeToFile(path string, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}
