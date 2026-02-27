package pocketlog_test

import (
	"fmt"
	"logger/pocketlog"
	"testing"
)

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name"
)

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level: pocketlog.LevelDebug,
			expected: "Debug - " + debugMessage + "\n" +
				"Info - " + infoMessage + "\n" +
				"Error - " + errorMessage + "\n",
		},
		"info": {
			level: pocketlog.LevelInfo,
			expected: "Info - " + infoMessage + "\n" +
				"Error - " + errorMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "Error - " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q",
					tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_Logf(t *testing.T) {
	type testcase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testcase{
		"Logf ErrorLevel": {
			level: pocketlog.LevelInfo,
			expected: "Info - " + infoMessage + "\n" +
				"Error - " + errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.Logf(pocketlog.LevelDebug, debugMessage)
			testedLogger.Logf(pocketlog.LevelInfo, infoMessage)
			testedLogger.Logf(pocketlog.LevelError, errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q",
					tc.expected, tw.contents)
			}
		})
	}
}

func TestLogger_WithLimit(t *testing.T) {
	const LIMIT = 10

	type testcase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testcase{
		fmt.Sprintf("Limit: %v", LIMIT): {
			level:    pocketlog.LevelError,
			expected: ("Error - " + errorMessage)[:LIMIT] + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw), pocketlog.WithLimit(LIMIT))

			testedLogger.Logf(pocketlog.LevelError, errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q",
					tc.expected, tw.contents)
			}
		})
	}
}
