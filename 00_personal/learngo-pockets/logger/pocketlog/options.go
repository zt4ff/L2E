package pocketlog

import "io"

type Option func(*Logger)

// WithOutput returns a configuration function that sets the outputs of logs
func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
