package gordle

import (
	"bufio"
	"fmt"
	"io"
)

// Game hold all the information needed to play the game Gordle
type Game struct {
	reader *bufio.Reader
}

// NewGame eturns a Game which can be used to play wht gordle game
func NewGame(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	return g
}

// Play starts the gordle game
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	fmt.Printf("Enter a guess:\n")
}
