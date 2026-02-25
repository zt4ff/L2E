package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	BooksRead []struct {
		BookID        string `json:"book_id"`
		Rating        int    `json:"rating"`
		DateRead      string `json:"date_read"`
		DidNotFinish  bool   `json:"did_not_finish"`
		ReviewSnippet string `json:"review_snippet"`
		Format        string `json:"format"`
		Reread        bool   `json:"reread"`
	} `json:"books_read"`
	Wishlist []struct {
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

func parseJson(path string) (BookDatabase, error) {
	f, err := os.Open(path)
	if err != nil {
		return BookDatabase{}, err
	}
	defer f.Close()

	var bookDatabase BookDatabase

	err = json.NewDecoder(f).Decode(&bookDatabase)

	if err != nil {
		return BookDatabase{}, err
	}

	return bookDatabase, nil
}

// ===================================================================

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}

type byAuthor []Book

// Len implements sort.Interface by returning the lenth of BookByAuthor
func (b byAuthor) Len() int { return len(b) }

func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements sort.Interface and returns books sorted by Author and then Title
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}

	return b[i].Title < b[j].Title
}
