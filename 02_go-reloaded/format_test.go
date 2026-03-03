package main

import (
	"reflect"
	"testing"
)

func TestFormatWord(t *testing.T) {
	testcases := map[string]struct {
		input    string
		expected string
	}{
		"test 1": {
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		"test 2": {
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		"test 3": {
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		"test 4": {
			input:    "Punctuation tests are ... kinda boring ,what do you think ?",
			expected: "Punctuation tests are... kinda boring, what do you think?",
		},
		"test 5": {
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
		"test 6": {
			input:    "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			expected: "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for name, testcase := range testcases {
		t.Run(name, func(t *testing.T) {
			got := formatWords(testcase.input)

			if got != testcase.expected {
				t.Errorf("expected: %s, got: %s", testcase.expected, got)
			}
		})
	}
}

func TestTokenize(t *testing.T) {
	tt := map[string]struct {
		input    string
		expected []string
	}{
		"simple command": {
			input:    "it (cap) was the best of times,",
			expected: []string{"it", "(cap)", "was", "the", "best", "of", "times", ","},
		},
		"no command word": {
			input:    "a happy function with no command",
			expected: []string{"a", "happy", "function", "with", "no", "command"},
		},
		"complex token": {
			input:    "foolishness (cap, 6) , it",
			expected: []string{"foolishness", "(cap, 6)", ",", "it"},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tokenizer(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tt := map[string]struct {
		input    string
		expected string
	}{
		"all lowercase": {
			input:    "hello",
			expected: "Hello",
		},
		"all uppercase": {
			input:    "HELLO",
			expected: "Hello",
		},
		"mixed case": {
			input:    "HEllO",
			expected: "Hello",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := capitalize(tc.input)
			if tc.expected != got {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestExtractNumber(t *testing.T) {
	tt := map[string]struct {
		input    string
		expected int
	}{
		"without numbers": {
			input:    "(cap)",
			expected: 1,
		},
		"with numbers": {
			input:    "(cap, 3)",
			expected: 3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := extractNumber(tc.input)
			if tc.expected != got {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestIsWord(t *testing.T) {
	tt := map[string]struct {
		input    string
		expected bool
	}{
		"normal word": {
			input:    "word",
			expected: true,
		},
		"no word": {
			input:    "",
			expected: false,
		},
		"wrong words": {
			input:    "(cap, 4)",
			expected: false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := isWord(tc.input)
			if tc.expected != got {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
