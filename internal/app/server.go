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

func (s *Server) FindBook(ctx context.Context, query *books.BookQuery) (*books.Book, error) {
	return s.repo.Find(query.Id)
}

func (s *Server) FindAllBooks(in *books.BookCollectionQuery, stream books.BookService_FindAllBooksServer) error {
	books, err := s.repo.FindAll()
	if err != nil {
		return err
	}
	for _, book := range books {
		stream.Send(book)
	}
	return nil
}

func (s *Server) AddBook(ctx context.Context, new *books.Book) (*books.Book, error) {
	return s.repo.Create(new)
}
