package app

import (
	"context"
	"github.com/benschw/books-grpc-poc/internal/fakes"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	lis := bufconn.Listen(1024 * 1024)

	s := grpc.NewServer()

	repo := fakes.NewRepo()

	books.RegisterBookServiceServer(s, NewServer(repo))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
}

func getConn(ctx context.Context) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func TestServer_AddBook(t *testing.T) {
	// given
	ctx := context.Background()
	conn := getConn(ctx)
	defer conn.Close()

	client := books.NewBookServiceClient(conn)

	newBook := &books.Book{Author: "Bob Loblaw", Title: "Law Blog"}

	// when
	createdBook, err := client.AddBook(ctx, newBook)

	// then
	assert.Nil(t, err)

	er, _ := status.FromError(err);
	assert.Equal(t, codes.OK, er.Code())

	assert.Equal(t, newBook.GetAuthor(), createdBook.GetAuthor())
	assert.Equal(t, newBook.GetTitle(), createdBook.GetTitle())
}

func TestServer_FindAllBook(t *testing.T) {
	// given
	ctx := context.Background()
	conn := getConn(ctx)
	defer conn.Close()

	client := books.NewBookServiceClient(conn)

	book1, err := client.AddBook(ctx, &books.Book{Author: "Bob Loblaw", Title: "Law Blog"})
	book2, err := client.AddBook(ctx, &books.Book{Author: "Bob Loblaw", Title: "Law Blog"})

	// when
	found, err := client.FindAllBooks(ctx, &books.BookQuery{})
	found1, err1 := found.Recv()
	found2, err2 := found.Recv()
	// then

	assert.Nil(t, err)
	assert.Nil(t, err1)
	assert.Nil(t, err2)

	er, _ := status.FromError(err);
	assert.Equal(t, codes.OK, er.Code())

	assert.Equal(t, book1.GetId(), found1.GetId())
	assert.Equal(t, book1.GetAuthor(), found1.GetAuthor())
	assert.Equal(t, book1.GetTitle(), found1.GetTitle())

	assert.Equal(t, book2.GetId(), found2.GetId())
	assert.Equal(t, book2.GetAuthor(), found2.GetAuthor())
	assert.Equal(t, book2.GetTitle(), found2.GetTitle())
}
