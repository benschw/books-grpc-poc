package main

import (
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
)




func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// ex
	c := books.NewBookServiceClient(conn)

	// add
	book, err := c.AddBook(context.Background(), &books.Book{Title: "Hello From Client!"})
	if err != nil {
		log.Fatalf("AddBook: Error when calling AddBook: %s", err)
	}
	log.Printf("AddBook: %v", book)

	// find
	book, err = c.FindBook(context.Background(), &books.BookQuery{Id: 1})
	if err != nil {
		log.Fatalf("FindBook: Error when calling FindBook: %s", err)
	}
	log.Printf("FindBook: %v", book)

	// findAll
	bookStream, err := c.FindAllBooks(context.Background(), &books.BookCollectionQuery{Author: "foo"})
	if err != nil {
		log.Fatalf("FindAllBooks: Error when calling FindAllBooks: %s", err)
	}
	for {
		book, err = bookStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("FindAllBooks: %s", err)
		}
		log.Printf("FindAllBooks: %v", book)
	}
}