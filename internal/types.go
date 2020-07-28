package internal

import "github.com/benschw/books-poc/models"

// Repo manages book persistence
type Repo interface {
	FindAll() ([]models.Book, error)
	Find(id uint64) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(id uint64) error
}
