package internal

import (
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
)

// Repo manages book persistence
type Repo interface {
	Create(book *books.Book) (*books.Book, error)
}
