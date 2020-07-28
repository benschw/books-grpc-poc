package main

import (
	"fmt"
	"os"

	"github.com/benschw/books-poc/internal/app"
	"github.com/benschw/books-poc/internal/postgres"
)

func main() {
	repo, err := postgres.NewRepo(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	app.NewWebServer(repo).Run()
}
