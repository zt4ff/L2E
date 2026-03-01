package main

import (
	"fmt"
	"logger/pocketlog"
	"os"
	"time"
)

func main() {
	filePath := "output.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	lgr := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(file))

	for _ = range 5 {
		lgr.Infof("A little copying is better than a little dependency.")
		lgr.Errorf("Errors are values. Documentation is for %s.", "users")
		lgr.Debugf("Make the zero (%d) value useful.", 0)
		lgr.Logf(pocketlog.LevelInfo, "Hallo, %d %v", 2022, time.Now())
	}

	// =====================
	// READ LOG OPERATIONS
	// =====================

	lgr.Reader.Stat()
}
