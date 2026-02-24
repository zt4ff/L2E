package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

// returns the books thar are on more than one bookworm's shelf
func findCommonBooks(bookworms []Bookworm) []Book {
	dict := make(map[string]uint)
	result := []Book{}
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			dict[book.Title]++

			if dict[book.Title] > 1 {
				result = append(result, book)
			}
		}
	}

	return result
}
