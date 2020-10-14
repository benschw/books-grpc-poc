package app

import (
	"github.com/benschw/books-grpc-poc/internal"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"golang.org/x/net/context"
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

