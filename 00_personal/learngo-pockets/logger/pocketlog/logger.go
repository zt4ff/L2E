package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the required threshold
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{
		threshold: threshold,
		output:    os.Stdout,
	}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// logf prints the message to the output.
func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format, args...)
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
