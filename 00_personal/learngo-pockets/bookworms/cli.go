package main

import (
	"flag"
	"fmt"
	"os"
)

type Mode string

const (
	ModeCommonBooks    = "cb"
	ModeRecommendation = "r"
)

func getCLIFlags() (string, string) {
	// CLI configurations
	var path string
	var mode string

	flag.StringVar(&path, "path", "testdata/bookworms_dataset.json", "the path to the database JSON")
	flag.StringVar(&path, "p", "testdata/bookworms_dataset.json", "the path to the database JSON")

	flag.StringVar(&mode, "m", "cb", "the mode of operations of the program. Options are 'cb' and 'r'")
	flag.StringVar(&mode, "mode", "cb", "the mode of operations of the program. Options are 'cb' and 'r'")

	flag.Parse()

	if _, err := os.Stat(path); err != nil {
		os.Exit(1)
		fmt.Fprintf(os.Stderr, "Error opening the file %s: %v", path, err)
		return "", ""
	}

	switch mode {
	case ModeCommonBooks:
	case ModeRecommendation:
		return mode, path
	default:
		fmt.Fprintf(os.Stderr, "You have a provided a wrong mode, the available options are '%s' and '%s'", ModeCommonBooks, ModeRecommendation)
		os.Exit(1)
	}

	return mode, path
}
