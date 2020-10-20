package app

import (
	"github.com/benschw/books-grpc-poc/internal"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"golang.org/x/net/context"
	"io"
)

type Server struct {
	repo internal.Repo
}

func NewServer(repo internal.Repo) *Server {
	return &Server{repo: repo}
}

func (s *Server) AddBook(ctx context.Context, new *books.Book) (*books.Book, error) {
	return s.repo.Create(new)
}

func (s *Server) FindAllBooks(query *books.BookQuery, stream books.BookService_FindAllBooksServer) error {
	books, err := s.repo.FindAll(query)
	if err != nil {
		return err
	}
	for _, book := range books {
		stream.Send(book)
	}
	return nil
}

func (s *Server) BulkAddBooks(stream books.BookService_BulkAddBooksServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		book, err := s.repo.Create(in)
		if err := stream.Send(book); err != nil {
			return err
		}
	}
}