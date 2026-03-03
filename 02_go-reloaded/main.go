package main

import (
	"fmt"
	"os"
)

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
	os.Exit(1)
}

func main() {
	files, err := parseArgument()
	if err != nil {
		printError(err)
	}

	input, err := readFile(files[0])
	if err != nil {
		printError(err)
	}

	formattedInput := formatWords(input)

	if err = writeToFile(files[1], formattedInput); err != nil {
		printError(err)
	}
}
