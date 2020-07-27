package main

import (
	"fmt"
	"os"

	"github.com/benschw/books-poc/app"
	"github.com/benschw/books-poc/postgres"
)

func main() {

	dbStr := os.Getenv("DATABASE_URL")

	repo, err := postgres.NewRepo(dbStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	webServer := app.NewWebServer(repo)
	webServer.Run()

}
