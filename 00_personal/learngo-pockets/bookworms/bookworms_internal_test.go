package main

import (
	"testing"
)

// create variables of reusable Books
var (
	handmaidsTale = Book{
		Author: "Margaret Atwood", Title: "The Handmaid's Tale",
	}
	oryxAndCrake = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar   = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre     = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

// func TestLoadBookworms_Success(t *testing.T) {
// 	tests := map[string]struct {
// 		bookwormsFile string
// 		want          []Bookworm
// 		wantErr       bool
// 	}{
// 		"file exists": {
// 			bookwormsFile: "testdata/bookworms.json",
// 			want: []Bookworm{
// 				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
// 				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
// 			},
// 			wantErr: false,
// 		},
// 		"file doesn't exist": {
// 			bookwormsFile: "testdata/bookworms.json",
// 			want:          nil,
// 			wantErr:       true,
// 		},
// 		"invalid JSON": {
// 			bookwormsFile: "testdata/invalid.json",
// 			want:          nil,
// 			wantErr:       true,
// 		},
// 	}

// 	for name, testCase := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			got, err := loadBookworms(testCase.bookwormsFile)
// 			if err != nil && !testCase.wantErr {
// 				t.Fatalf("expected an error %s, got none", err)
// 			}

// 			if err == nil && testCase.wantErr {
// 				t.Fatalf("expected no error, got one %s", err)
// 			}

// 			if !reflect.DeepEqual(got, testCase.want) {
// 				t.Fatalf("diffterent result: got %v, expected %v", got, testCase.want)

// 			}
// 		})
// 	}
// }

// ===============================================

func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	if len(got) != len(want) {
		return false
	}

	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}

	return true
}

func TestBooksCount(t *testing.T) {
	tt := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{
					handmaidsTale, theBellJar,
				}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{
					handmaidsTale, theBellJar, theBellJar,
				}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    2,
				oryxAndCrake:  1,
				janeEyre:      1},
		},
	}

	// running the test
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := booksCount(tc.input)
			if !equalBooksCount(t, tc.want, got) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}
