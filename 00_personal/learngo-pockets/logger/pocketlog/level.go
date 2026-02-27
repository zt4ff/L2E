package pocketlog

// Level represents an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains informated deemed inportant
	LevelInfo
	// LevelError represents a logging level for errors
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelError:
		return "Error"
	default:
		return "Unknown"
	}
}
