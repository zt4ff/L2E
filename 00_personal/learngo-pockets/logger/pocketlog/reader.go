package pocketlog

import (
	"bufio"
	"fmt"
	"os"
)

// Reader is used to read log information
type Reader struct {
	logger *Logger
}

func newReader(logger *Logger) *Reader {
	return &Reader{logger: logger}
}

func (r *Reader) openFile() (*os.File, error) {
	f, ok := r.logger.output.(*os.File)
	if !ok {
		return nil, fmt.Errorf("logger output is not a file")
	}
	return os.Open(f.Name())
}

func (r *Reader) Head(lines int) {
	file, err := r.openFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < lines && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading log: %v\n", err)
	}
}

func (r *Reader) Tail(lines int) {
	file, err := r.openFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Ring buffer: only keep the last `lines` lines in memory
	buf := make([]string, lines)
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buf[count%lines] = scanner.Text()
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading log: %v\n", err)
		return
	}

	start, total := 0, min(count, lines)
	if count > lines {
		start = count % lines
	}
	for i := range total {
		fmt.Println(buf[(start+i)%lines])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
