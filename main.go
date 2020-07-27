package main

import (
	"context"
	"fmt"
	"os"

	"github.com/benschw/books-poc/data"
	"github.com/benschw/books-poc/resources"
	"github.com/gin-gonic/gin"

	"github.com/jackc/pgx/v4"
)

func main() {

	//dbURL := os.Getenv("DATABASE_URL")
	dbURL := "postgres://docker:docker@localhost:5400/books"

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	repo := &data.BooksRepo{
		Conn: conn,
	}
	resource := &resources.BooksResource{
		Repo: repo,
	}

	// Routes
	r.GET("/book", resource.FindBooks)
	r.GET("/book/:id", resource.FindBook)
	r.POST("/book", resource.CreateBook)
	r.PUT("/book/:id", resource.UpdateBook)
	r.DELETE("/book/:id", resource.DeleteBook)

	r.Run()
}
