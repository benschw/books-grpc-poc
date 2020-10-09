package fakes

import (
	"fmt"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"

	"github.com/benschw/books-grpc-poc/internal"
)

// Ensure Repo implements internal.BooksRepo.
var _ internal.Repo = &Repo{}

// Repo manages fake database access for books_old
type Repo struct {
	i     uint64
	Books []*books.Book
}

// NewRepo creates a new postgres repo
func NewRepo() *Repo {

	return &Repo{i: 0, Books: []*books.Book{}}
}

// FindAll returns all books_old from the database
func (r *Repo) FindAll() ([]*books.Book, error) {
	return r.Books, nil
}

// Find selects one book by id from the database
func (r *Repo) Find(id uint64) (*books.Book, error) {
	var book *books.Book

	for _, b := range r.Books {
		if b.Id == id {
			return b, nil
		}
	}
	return book, fmt.Errorf("not found")
}

// Create adds a new book to the databases
func (r *Repo) Create(book *books.Book) (*books.Book, error) {
	r.i = r.i + 1
	book.Id = r.i
	r.Books = append(r.Books, book)
	return book, nil
}

// Update updates an existing record in the database
func (r *Repo) Update(book *books.Book) (*books.Book, error) {
	for i, b := range r.Books {
		if b.Id == book.Id {
			r.Books[i].Title = book.Title
			r.Books[i].Author = book.Author
			return book, nil
		}
	}

	return book, fmt.Errorf("not found")
}

// Delete deletes a record in the database
func (r *Repo) Delete(id uint64) error {

	for i, b := range r.Books {
		if b.Id == id {
			// Remove the element at index i from a.
			copy(r.Books[i:], r.Books[i+1:])   // Shift a[i+1:] left one index.
			r.Books = r.Books[:len(r.Books)-1] // Truncate slice.
			return nil
		}
	}

	return fmt.Errorf("not found")
}
