package internal

import "github.com/benschw/books-poc/books"

// Repo manages book persistence
type Repo interface {
	FindAll() ([]books.Book, error)
	Find(id uint64) (books.Book, error)
	Create(book books.Book) (books.Book, error)
	Update(book books.Book) (books.Book, error)
	Delete(id uint64) error
}
