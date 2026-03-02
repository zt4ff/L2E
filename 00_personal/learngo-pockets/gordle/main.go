package main

import (
	"gordle/gordle"
	"os"
)

const maxAttempts = 6

func main() {
	solution := "hello"
	game := gordle.NewGame(os.Stdin, solution, maxAttempts)

	game.Play()
}
