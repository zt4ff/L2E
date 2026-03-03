package main

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var vowels = []string{"a", "e", "i", "o", "u", "h"}

// tokenizer split all entity in a string into a slice of words
func tokenizer(str string) []string {
	var tokens []string
	var current strings.Builder
	inParen := false

	for _, r := range str {
		switch {
		case r == '(':
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			inParen = true
			current.WriteRune(r)

		case r == ')':
			current.WriteRune(r)
			tokens = append(tokens, current.String())
			current.Reset()
			inParen = false

		case inParen:
			current.WriteRune(r)

		case unicode.IsSpace(r):
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}

		case strings.ContainsRune(".,!?:;'", r):
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(r))

		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

// capitalize simple capitilize the strings provided to it.
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// extractNumber extracts the number of a command - (<cmd>) and return 1 if not any
func extractNumber(cmd string) int {
	cmd = strings.TrimSuffix(cmd, ")")
	parts := strings.Split(cmd, ",")
	if len(parts) != 2 {
		return 1
	}
	n, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	return n
}

// a word is one without the command or any puntuation
func isWord(token string) bool {
	if len(token) < 1 {
		return false
	}
	return !strings.HasPrefix(token, "(") &&
		!strings.ContainsAny(token, ".,!?:;'")
}

// the word in token is always a command,
func applyCommand(tokens []string, i int) ([]string, int) {
	cmd := tokens[i]

	applyFnOnPrevWords := func(n int, fn func(int)) {
		count := 0
		for j := i - 1; j >= 0 && count < n; j-- {
			if isWord(tokens[j]) {
				fn(j)
				count++
			}
		}
	}

	switch {
	case cmd == "(hex)":
		n, _ := strconv.ParseInt(tokens[i-1], 16, 64)
		tokens[i-1] = strconv.FormatInt(n, 10)

	case cmd == "(bin)":
		n, _ := strconv.ParseInt(tokens[i-1], 2, 64)
		tokens[i-1] = strconv.FormatInt(n, 10)

	case cmd == "(up)":
		applyFnOnPrevWords(1, func(j int) {
			tokens[j] = strings.ToUpper(tokens[j])
		})

	case cmd == "(low)":
		applyFnOnPrevWords(1, func(j int) {
			tokens[j] = strings.ToLower(tokens[j])
		})

	case cmd == "(cap)":
		applyFnOnPrevWords(1, func(j int) {
			tokens[j] = capitalize(tokens[j])
		})

	case strings.HasPrefix(cmd, "(up,"):
		applyFnOnPrevWords(extractNumber(cmd), func(j int) {
			tokens[j] = strings.ToUpper(tokens[j])
		})

	case strings.HasPrefix(cmd, "(low,"):
		applyFnOnPrevWords(extractNumber(cmd), func(j int) {
			tokens[j] = strings.ToLower(tokens[j])
		})

	case strings.HasPrefix(cmd, "(cap,"):
		applyFnOnPrevWords(extractNumber(cmd), func(j int) {
			tokens[j] = capitalize(tokens[j])
		})
	}

	tokens = append(tokens[:i], tokens[i+1:]...)
	return tokens, i - 1
}

func fixAAndExc(tokens []string) {
	for i := 0; i < len(tokens)-1; i++ {
		if strings.ToLower(tokens[i]) == "a" && isWord(tokens[i+1]) {
			first := strings.ToLower(string(tokens[i+1][0]))
			if slices.Contains(vowels, first) {
				if tokens[i] == "A" {
					tokens[i] = "An"
				} else {
					tokens[i] = "an"
				}
			}
		}
	}
}

func rebuild(tokens []string) string {
	var sb strings.Builder
	inQuote := false

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "'" {
			if !inQuote {
				if i > 0 {
					sb.WriteString(" ")
				}
				sb.WriteString("'")
				inQuote = true
			} else {
				sb.WriteString("'")
				inQuote = false
			}
			continue
		}

		if i > 0 && !strings.ContainsAny(tokens[i], ".,!?:;") {

			// in quote and last item of a string builder is '
			if inQuote && sb.String()[sb.Len()-1] == '\'' {
			} else {
				sb.WriteString(" ")
			}
		}

		sb.WriteString(tokens[i])
	}

	return sb.String()
}

func formatWords(str string) string {
	tokens := tokenizer(str)

	for i := 0; i < len(tokens); i++ {
		if strings.HasPrefix(tokens[i], "(") {
			tokens, i = applyCommand(tokens, i)
		}
	}

	fixAAndExc(tokens)

	return rebuild(tokens)
}
