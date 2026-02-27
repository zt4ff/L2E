package pocketlog

// Level represents an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains informated deemed inportant
	LevelInfo
	// LevelWarn represents a logging level to warn users
	LevelWarn
	// LevelError represents a logging level for errors
	LevelError
	// LevelError represents a logging level that indicates fatal situation in the system
	LevelFatal
)
