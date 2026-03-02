package gordle

import (
	"errors"
	"reflect"
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
			g, _ := NewGame(strings.NewReader(tc.input), []string{}, 0)

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
			g, _ := NewGame(nil, []string{""}, 0)

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

func TestFeedString(t *testing.T) {
	tt := map[string]struct {
		input    feedback
		expected string
	}{
		"nominal": {
			input:    feedback{absentCharacter, wrongPosition, correctPosition},
			expected: "⬜🟡💚",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.input.String()
			if got != tc.expected {
				t.Errorf("expected: %s, got :%s", tc.expected, got)
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess    []rune
		solution []rune
		expected string
	}{
		"nominal": {
			guess:    toUppercaseCharacters("small"),
			solution: toUppercaseCharacters("hello"),
			expected: "⬜⬜⬜💚🟡",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback(tc.guess, tc.solution)

			got := fb.String()
			if got != tc.expected {
				t.Errorf("expected: %s, got %s", tc.expected, got)
			}
		})
	}
}

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		path    string
		wantErr bool
		expect  []string
	}{
		"nominal": {
			path:    "testdata/test_text.txt",
			wantErr: false,
			expect:  []string{"apple", "brave", "chair", "dream", "eagle"},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ReadCorpus(tc.path)
			if !tc.wantErr && err != nil {
				t.Errorf("got an error when it's not expected: %v", err)
			}

			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("expecting: %s, got: %s", tc.expect, got)
			}
		})

	}
}
