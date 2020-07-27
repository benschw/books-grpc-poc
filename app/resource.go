package app

import (
	"net/http"
	"strconv"

	"github.com/benschw/books-poc"
	"github.com/gin-gonic/gin"
)

// BooksResource provides handlers for the Books Resource
type BooksResource struct {
	repo books.Repo
}

// FindBooks finds all books
// GET /books
func (r *BooksResource) FindBooks(c *gin.Context) {
	books, err := r.repo.FindAll()
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

	book, err := r.repo.Find(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook adds a new book
// POST /books
func (r *BooksResource) CreateBook(c *gin.Context) {
	var input books.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := r.repo.Create(input)
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

	found, err := r.repo.Find(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input books.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = found.ID

	updated, err := r.repo.Update(input)
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

	_, err := r.repo.Find(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = r.repo.Delete(i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": nil})
}
