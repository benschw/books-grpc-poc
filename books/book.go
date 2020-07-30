package books

// Book model for JSON api
type Book struct {
	ID     uint64 `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
