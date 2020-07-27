package resources

import (
	"net/http"
	"strconv"

	"github.com/benschw/books-poc/api"
	"github.com/benschw/books-poc/data"
	"github.com/gin-gonic/gin"
)

// BooksResource provides handlers for the Books Resource
type BooksResource struct {
	Repo *data.BooksRepo
}

// FindBooks finds all books
// GET /books
func (r *BooksResource) FindBooks(c *gin.Context) {
	books, err := r.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// FindBook finds a book by id
// GET /books/:id
func (r *BooksResource) FindBook(c *gin.Context) {
	i, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	book, err := r.Repo.Find(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook adds a new book
// POST /books
func (r *BooksResource) CreateBook(c *gin.Context) {
	var input api.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := r.Repo.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": created})
}

// UpdateBook updates an existing book
// PUT /books/:id
func (r *BooksResource) UpdateBook(c *gin.Context) {
	i, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	found, err := r.Repo.Find(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input api.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = found.ID

	updated, err := r.Repo.Update(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updated})
}

// DeleteBook deletes a book
// DELETE /books/:id
func (r *BooksResource) DeleteBook(c *gin.Context) {
	i, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	_, err := r.Repo.Find(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = r.Repo.Delete(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": nil})
}
