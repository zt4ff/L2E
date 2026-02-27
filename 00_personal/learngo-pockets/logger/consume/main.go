package main

import (
	"fmt"
	"logger/pocketlog"
)

func main() {
	logger := pocketlog.New(pocketlog.LevelError)
	fmt.Println(logger)
}
