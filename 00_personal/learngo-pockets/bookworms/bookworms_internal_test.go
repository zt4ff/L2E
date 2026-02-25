package main

import (
	"reflect"
	"testing"
)

// create variables of reusable Books
var (
	handmaidsTale = Book{
		BookID:    "b001",
		Title:     "The Handmaid's Tale",
		Author:    "Margaret Atwood",
		AuthorID:  "a001",
		Genres:    []string{"dystopia", "fiction", "feminist"},
		Year:      1985,
		Pages:     311,
		AvgRating: 4.1,
		Language:  "en",
		Isbn:      "9780385490818",
	}

	orxyAndCrake = Book{
		BookID:   "b002",
		Title:    "Oryx and Crake",
		Author:   "Margaret Atwood",
		AuthorID: "a001",
		Genres: []string{"sci-fi",
			"dystopia",
			"fiction"},
		Year:      2003,
		Pages:     374,
		AvgRating: 4.0,
		Language:  "en",
		Isbn:      "9780385721677",
	}

	theBellJar = Book{
		BookID:   "b003",
		Title:    "The Bell Jar",
		Author:   "Sylvia Plath",
		AuthorID: "a002",
		Genres: []string{"fiction",
			"classic",
			"mental-health"},
		Year:      1963,
		Pages:     244,
		AvgRating: 4.0,
		Language:  "en",
		Isbn:      "9780060837020",
	}

	janeEyre = Book{
		BookID:   "b004",
		Title:    "Jane Eyre",
		Author:   "Charlotte Bront\u00eb",
		AuthorID: "a003",
		Genres: []string{"classic",
			"romance",
			"fiction"},
		Year:      1847,
		Pages:     532,
		AvgRating: 4.2,
		Language:  "en",
		Isbn:      "9780141441146",
	}
)

func TestParsingJson(t *testing.T) {
	testcases := map[string]struct {
		input   string // file path
		want    []Book
		wantErr bool
	}{
		"get all books": {
			input:   "testdata/bookworms.json",
			want:    []Book{handmaidsTale, orxyAndCrake, theBellJar, janeEyre},
			wantErr: false,
		},
	}

	for name, testcase := range testcases {
		t.Run(name, func(t *testing.T) {
			bookDatabase, err := parseJson(testcase.input)

			if testcase.wantErr && err == nil {
				t.Fatalf("expected error but got no error")
			}

			if !testcase.wantErr && err != nil {
				t.Fatalf("expected no error but got error: %v", err)
			}

			if !reflect.DeepEqual(bookDatabase.Books, testcase.want) {
				t.Fatalf("books don't match, expect %v but got %v", testcase.want, bookDatabase.Books)
			}
		})
	}
}
