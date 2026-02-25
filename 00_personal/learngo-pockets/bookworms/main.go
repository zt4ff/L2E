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

	// find common book

	for key, _ := range bookDatabase.Bookworms {
		fmt.Println(key)
	}
}
