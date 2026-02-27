package pocketlog

import "fmt"

// Logger is used to log information
type Logger struct {
	threshold Level
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Infof(format string, args ...any) {
	// todo
}

func (l *Logger) Errorf(format string, args ...any) {
	// todo
}

// New returns you a logger, ready to log at the required threshold
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
