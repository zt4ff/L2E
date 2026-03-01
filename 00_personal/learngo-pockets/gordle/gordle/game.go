package gordle

import "fmt"

// Game hold all the information needed to play the game Gordle
type Game struct {

}

// NewGame eturns a Game which can be used to play wht gordle game
func NewGame() *Game {
	g := &Game{}

	return g
}


// Play starts the gordle game
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle")

	fmt.Printf("Enter a guess:\n")
}