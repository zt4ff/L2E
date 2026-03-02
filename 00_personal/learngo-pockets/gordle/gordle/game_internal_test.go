package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english lowercase": {
			input: "hello",
			want:  []rune("HELLO"),
		},
		"5 characters in english uppercase": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := NewGame(strings.NewReader(tc.input), "", 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %s, want = %s", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidation(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("POCKET"),
			expected: errInvalidWordLength,
		},
		"too short": {
			word:     []rune("SHO"),
			expected: errInvalidWordLength,
		},
		"nil character": {
			word:     []rune{},
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := NewGame(nil, "", 0)

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestToUpperCase(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"all lowercase": {
			input: "welcome",
			want:  []rune("WELCOME"),
		},
		"all uppercase": {
			input: "JOHN",
			want:  []rune("JOHN"),
		},
		"mixed cases": {
			input: "wElCoME",
			want:  []rune("WELCOME"),
		},
		"nil input": {
			input: "",
			want:  []rune{},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := toUppercaseCharacters(tc.input)

			if !slices.Equal(got, tc.want) {
				t.Errorf("expected: %v, got %v", tc.want, got)
			}
		})
	}
}
