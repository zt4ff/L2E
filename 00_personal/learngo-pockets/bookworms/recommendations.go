package main

type set map[Book]struct{}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

func listOtherBooksOnShelves(i int, books []Book) []Book {
	var filteredBooks []Book

	for x, book := range books {
		if x != i {
			filteredBooks = append(filteredBooks, book)
		}
	}

	return filteredBooks
}

// func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
// 	sb := make(bookRecommendations)

// 	// register all books on everyone's shelf.
// 	for _, bookworm := range bookworms {
// 		for i, book := range bookworm.Books {
// 			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
// 			registerBookRecommendations(sb, book, otherBooksOnShelves)
// 		}
// 	}

// 	// Recommend a lisf of related books to each bookworm
// 	recommendations := make([]Bookworm, len(bookworms))
// 	for i, bookworm := range bookworms {
// 		recommendations[i] = Bookworm{
// 			Name:  bookworm.Name,
// 			Books: recommendBooks(sb, bookworm.Books),
// 		}
// 	}

// 	return recommendations
// }
