package pocketlog

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reader is used to read log information
type Reader struct {
	logger *Logger
}

func newReader(logger *Logger) *Reader {
	return &Reader{logger: logger}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
}

func getLevelAndMessage(line string) (level string, msg string) {
	split := strings.Split(line, " - ")

	return split[0], split[1]
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
		printError(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < lines && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		printError(err)
	}
}

func (r *Reader) Tail(lines int) {
	file, err := r.openFile()
	if err != nil {
		printError(err)
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

func (r *Reader) Stat() map[string]uint {
	file, err := r.openFile()
	if err != nil {
		printError(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stat := make(map[string]uint)

	for scanner.Scan() {
		line := scanner.Text()
		level, _ := getLevelAndMessage(line)

		stat[level]++
	}

	return stat
}

func (r *Reader) PrintAll() {
	file, err := r.openFile()
	if err != nil {
		printError(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
