package app

import (
	"github.com/benschw/books-poc/internal"
	"github.com/gin-gonic/gin"
)

// WebServer for managing books using gin
type WebServer struct {
	router *gin.Engine
}

// NewWebServer creates a new WebServer
func NewWebServer(repo internal.Repo) *WebServer {
	resource := &BooksResource{repo: repo}

	r := gin.Default()

	// Routes
	r.GET("/book", resource.FindBooks)
	r.GET("/book/:id", resource.FindBook)
	r.POST("/book", resource.CreateBook)
	r.PUT("/book/:id", resource.UpdateBook)
	r.DELETE("/book/:id", resource.DeleteBook)

	return &WebServer{
		router: r,
	}
}

// Router returns the configured Gin Router (Engine)
func (w *WebServer) Router() *gin.Engine {
	return w.router
}

// Run runs the web server
func (w *WebServer) Run() {
	w.Router().Run()
}
