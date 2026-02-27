package main

import (
	"fmt"
	"os"
)

func main() {
	mode, path := getCLIFlags()

	if mode == ModeCommonBooks {
		bookDatabase, err := parseJSON[BookDatabase](path)
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

	if mode == ModeRecommendation {
		recommendationDatabase, err := parseJSON[RecommendationsDatabase](path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing JSON: %v", err)
		}

		fmt.Println(recommendationDatabase)
	}
}
