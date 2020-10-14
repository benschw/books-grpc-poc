package main

import (
	"flag"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
)


var (
	addr = flag.String("addr", "localhost:9000", "The server address in the format of host:port")
	cmd = flag.String("cmd", "", "The client command to run (add or list)")
	author = flag.String("author", "", "author value for adding a book or searching by author")
	title = flag.String("title", "", "title value for adding a book")
)

func Add(c books.BookServiceClient, author string, title string) error {
	book, err := c.AddBook(context.Background(), &books.Book{Author: author, Title: title})
	if err != nil {
		return err
	}
	log.Printf("AddBook: %v", book)
	return nil
}

func List(c books.BookServiceClient, author string) error {
	bookStream, err := c.FindAllBooks(context.Background(), &books.BookQuery{Author: author})
	if err != nil {
		return err
	}
	for {
		book, err := bookStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("FindAllBooks: %v", book)
	}
	return nil
}

func main() {

	flag.Parse()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := books.NewBookServiceClient(conn)

	switch *cmd {
	case "add":
		if err := Add(c, *author, *title); err != nil {
			log.Fatalf("add - error adding book: %s", err)
		}
		break;
	case "list":
		if err := List(c, *author); err != nil {
			log.Fatalf("list - error listing books: %s", err)
		}
		break;
	default:
		log.Fatalf("unknown command: %s", *cmd)
	}
}