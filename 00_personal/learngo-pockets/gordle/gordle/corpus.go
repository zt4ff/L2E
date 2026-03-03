package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// corpusError defines a sentinel error
type corpusError string

func (e corpusError) Error() string {
	return string(e)
}

const ErrCorpusIsEmpty = corpusError("corpus is empty")

func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	words := strings.Fields(string(data))

	return words, nil
}

// pickWord returns a random word from the corpus
func pickWord(corpus []string) string {
	index := rand.Intn(len(corpus))

	return corpus[index]
}
