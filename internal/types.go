package internal

import (
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
)

// Repo manages book persistence
type Repo interface {
	FindAll(query *books.BookQuery) ([]*books.Book, error)
	Create(book *books.Book) (*books.Book, error)
}
