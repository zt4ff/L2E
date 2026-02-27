package main

import (
	"fmt"
	"logger/pocketlog"
	"os"
)

func main() {
	logger := pocketlog.New(pocketlog.LevelError, os.Stdout)
	fmt.Println(logger)
}
