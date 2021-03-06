package fakes

import (
	"github.com/benschw/books-grpc-poc/pkg/pb/books"

	"github.com/benschw/books-grpc-poc/internal"
)

// Ensure Repo implements internal.BooksRepo.
var _ internal.Repo = &Repo{}

// Repo manages fake database access for books
type Repo struct {
	i     uint64
	Books []*books.Book
}

// NewRepo creates a new in memory repo (create an implementation backed by a db for persistence)
func NewRepo() *Repo {

	return &Repo{i: 0, Books: []*books.Book{}}
}

// FindAll returns all books from the database, optionally filtered by author
func (r *Repo) FindAll(query *books.BookQuery) ([]*books.Book, error) {
	if query.Author == "" {
		return r.Books, nil
	}
	matches := []*books.Book{}
	for _, book := range r.Books {
		if book.Author == query.Author {
			matches = append(matches, book)
		}
	}
	return matches, nil
}

// Create adds a new book to the database
func (r *Repo) Create(book *books.Book) (*books.Book, error) {
	r.i = r.i + 1
	book.Id = r.i
	r.Books = append(r.Books, book)
	return book, nil
}
