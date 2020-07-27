package books

// Book model for JSON api
type Book struct {
	ID     uint64 `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Repo manages book persistence
type Repo interface {
	FindAll() ([]Book, error)
	Find(id uint64) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(id uint64) error
}

// WebServer runs a REST service API to interact with books
type WebServer interface {
	Run()
}
