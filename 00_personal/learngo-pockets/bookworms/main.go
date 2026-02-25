package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("filepath", "testdata/bookworms.json", "path to a  JSON file")

	flag.Parse()

	bookworms, err := loadBookworms(*filePath)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms %s\n", err)
		os.Exit(1)
	}

	displayBooks(findCommonBooks(bookworms))
}
