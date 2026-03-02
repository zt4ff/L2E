package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// ====================
// Helpers
// ====================

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
	reader      *bufio.Reader
	solution    []rune
	maxAttempts uint
}

// NewGame eturns a Game which can be used to play wht gordle game
func NewGame(playerInput io.Reader, solution string, maxAttempts uint) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    toUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return g
}

// Play starts the gordle game
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := uint(1); currentAttempt <= (g.maxAttempts); currentAttempt++ {
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("You've lost! The solution was: %s. \n", string(g.solution))

	// fmt.Printf("Enter a guess:\n")
}

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

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
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}
