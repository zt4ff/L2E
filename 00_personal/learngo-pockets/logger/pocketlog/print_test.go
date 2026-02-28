package pocketlog_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// testing output made to the Stdout

func PrintSomething(t *testing.T, str string) {
	t.Helper()
	fmt.Printf("Print something")
}

func Test_PrintSomething(t *testing.T) {

	type testcase struct {
		input string
		want  string
	}

	tt := map[string]testcase{
		"happy path": {
			input: "testing",
			want:  "testing",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			PrintSomething(t, tc.input)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			buf.ReadFrom(r)

			got := buf.String()

			if got != tc.want {
				t.Errorf("got: '%s' while expecting: '%s'", got, tc.want)
			}

		})
	}

}
