package main

import (
	"fmt"
)

type Book struct {
	BookID    string   `json:"book_id"`
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	AuthorID  string   `json:"author_id"`
	Genres    []string `json:"genres"`
	Year      int      `json:"year"`
	Pages     int      `json:"pages"`
	AvgRating float64  `json:"avg_rating"`
	Language  string   `json:"language"`
	Isbn      string   `json:"isbn"`
}

type BookRead struct {
	BookID        string `json:"book_id"`
	Rating        int    `json:"rating"`
	DateRead      string `json:"date_read"`
	DidNotFinish  bool   `json:"did_not_finish"`
	ReviewSnippet string `json:"review_snippet"`
	Format        string `json:"format"`
	Reread        bool   `json:"reread"`
}

type Bookworm struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location struct {
		City     string `json:"city"`
		Country  string `json:"country"`
		Timezone string `json:"timezone"`
	} `json:"location"`
	MemberSince string `json:"member_since"`
	Preferences struct {
		FavoriteGenres  []string `json:"favorite_genres"`
		DislikedGenres  []string `json:"disliked_genres"`
		FavoriteAuthors []string `json:"favorite_authors"`
		ReadingPace     string   `json:"reading_pace"`
		PreferredMoods  []string `json:"preferred_moods"`
		PreferredLength string   `json:"preferred_length"`
		PreferredEra    []string `json:"preferred_era"`
	} `json:"preferences"`
	ReadingStats struct {
		TotalBooksRead     int     `json:"total_books_read"`
		ReadingGoalPerYear int     `json:"reading_goal_per_year"`
		BooksReadThisYear  int     `json:"books_read_this_year"`
		AvgRatingGiven     float64 `json:"avg_rating_given"`
		LongestStreakDays  int     `json:"longest_streak_days"`
		MostReadGenre      string  `json:"most_read_genre"`
		MostReadAuthorID   string  `json:"most_read_author_id"`
	} `json:"reading_stats"`
	BooksRead []BookRead `json:"books_read"`
	Wishlist  []struct {
		BookID    string `json:"book_id"`
		AddedDate string `json:"added_date"`
		Priority  string `json:"priority"`
	} `json:"wishlist"`
}

type BookDatabase struct {
	SchemaVersion string `json:"schema_version"`
	Description   string `json:"description"`
	ModelHints    struct {
		Target              string   `json:"target"`
		FeaturesToEngineer  []string `json:"features_to_engineer"`
		SuggestedAlgorithms []string `json:"suggested_algorithms"`
	} `json:"model_hints"`
	Books     []Book     `json:"books"`
	Bookworms []Bookworm `json:"bookworms"`
}

// gets a struct of books read with the number of their appearance across bookworms
func getBooksDir(bookworms []Bookworm) map[string]int {
	booksDir := make(map[string]int)

	for _, bookworm := range bookworms {
		for _, bookRead := range bookworm.BooksRead {
			booksDir[bookRead.BookID]++
		}
	}

	return booksDir
}

func getBookbyId(id string, bookDatabase BookDatabase) (Book, bool) {
	for _, book := range bookDatabase.Books {
		if id == book.BookID {
			return book, true
		}
	}
	return Book{}, false
}

func displayBooks(books []Book) {
	fmt.Print("Common books found with bookworms \n\n")
	for _, book := range books {
		fmt.Printf("Book - %s by %s\n", book.Title, book.Author)
	}
}
