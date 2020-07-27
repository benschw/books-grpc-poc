package app

import (
	"github.com/benschw/books-poc"
	"github.com/gin-gonic/gin"
)

// Ensure WebServer implements books.WebServer.
var _ books.WebServer = &WebServer{}

// WebServer for managing books using gin
type WebServer struct {
	resource *BooksResource
}

// NewWebServer creates a new WebServer
func NewWebServer(repo books.Repo) *WebServer {
	return &WebServer{
		resource: &BooksResource{repo: repo},
	}
}

// Run starts the WebServer
func (w *WebServer) Run() {
	r := gin.Default()

	// Routes
	r.GET("/book", w.resource.FindBooks)
	r.GET("/book/:id", w.resource.FindBook)
	r.POST("/book", w.resource.CreateBook)
	r.PUT("/book/:id", w.resource.UpdateBook)
	r.DELETE("/book/:id", w.resource.DeleteBook)

	r.Run()

}
