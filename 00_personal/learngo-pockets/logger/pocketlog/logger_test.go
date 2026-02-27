package pocketlog_test

import (
	"logger/pocketlog"
	"os"
)

// Go won't allow use have more than one package in a folder. But if we surfix the package name with
// _test, it's allowed. And you'd have to import pocketlog package differently here

// Doing this is to allow use be able to to write closed-box texting where we are testing from
// outside of the package

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))
	debugLogger.Debugf("Hello, %s", "world")
}
