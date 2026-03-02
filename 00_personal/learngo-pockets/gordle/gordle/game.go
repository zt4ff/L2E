package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ====================
// Helpers
// ====================

const solutionLength = 5

// errInvalidWordLength is returned when the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// toUppercaseCharacters is a native implementtion to turn a string to Uppercase
func toUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// ====================
// Logic
// ====================

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
	guess := g.ask()

	fmt.Printf("Your guess is: %s\n", string(guess))
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := toUppercaseCharacters(string(playerInput))
		if err := g.validateGuess(guess); err != nil {
			_, _ = fmt.Printf("Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

// validateGuess ensures the guess is valid
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}

	return nil
}
