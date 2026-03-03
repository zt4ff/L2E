package main

import (
	"fmt"
	"gordle/gordle"
	"os"
)

const maxAttempts = 6

type customError string

func (e customError) Error() string {
	return string(e)
}

const CustomError = customError("customer error message")

func main() {
	corpus, err := gordle.ReadCorpus("testdata/text.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	game, err := gordle.NewGame(os.Stdin, corpus, maxAttempts)
	if err != nil {
	}

	game.Play()
}
