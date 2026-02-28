package pocketlog

import "fmt"

// Reader is used to read log information
type Reader struct {
	logger *Logger
}

func newReader(logger *Logger) *Reader {
	reader := Reader{
		logger: logger,
	}
	return &reader
}

func (r *Reader) Head(lines int) {
	// for i := 0; i < lines; i++ {
	//
	// }

	fmt.Println(r.logger.output)
}
