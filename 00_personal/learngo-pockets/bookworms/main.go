package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	path := flag.String("path", "testdata/bookworms_dataset.json", "the path to the database JSON")
	flag.Parse()

	bookDatabase, err := parseJson(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing JSON: %v", err)
	}

	booksDir := getBooksDir(bookDatabase.Bookworms)

	var books []Book

	for bookId, count := range booksDir {
		if count > 1 {
			book, ok := getBookbyId(bookId, bookDatabase)
			if ok {
				books = append(books, book)
			}
		}
	}

	displayBooks(books)
}
