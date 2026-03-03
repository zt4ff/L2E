package main

import (
	"gordle/gordle"
	"os"
)

func main() {
	game := gordle.NewGame(os.Stdin)

	game.Play()
}