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
	limit     uint
	Reader    Reader
}

// New returns you a logger, ready to log at the required threshold
func New(threshold Level, opts ...Option) *Logger {
	lgr := Logger{
		threshold: threshold,
		output:    os.Stdout,
		limit:     1000, // default
	}

	lgr.Reader = *newReader(&lgr)

	for _, configFunc := range opts {
		configFunc(&lgr)
	}

	return &lgr
}

func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	msg := fmt.Sprintf("%s - "+format, append([]any{lvl}, args...)...)

	if uint(len(msg)) > l.limit {
		msg = msg[:l.limit]
	}

	fmt.Fprintln(l.output, msg)
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	l.Logf(LevelDebug, format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	l.Logf(LevelInfo, format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	l.Logf(LevelError, format, args...)
}
