package postgres

import (
	"context"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"

	"github.com/benschw/books-grpc-poc/internal"
	"github.com/jackc/pgx/v4"
)

// Ensure Repo implements internal.BooksRepo.
var _ internal.Repo = &Repo{}

// Repo manages postgres database access for books_old
type Repo struct {
	conn *pgx.Conn
}

// NewRepo creates a new postgres repo
func NewRepo(dbStr string) (*Repo, error) {

	conn, err := pgx.Connect(context.Background(), dbStr)

	return &Repo{conn: conn}, err
}

// FindAll returns all books_old from the database
func (r *Repo) FindAll() ([]*books.Book, error) {
	var bs []*books.Book

	rows, _ := r.conn.Query(context.Background(), "select * from books_old order by id")

	for rows.Next() {
		var id uint64
		var title string
		var author string
		err := rows.Scan(&id, &title, &author)
		if err != nil {
			return bs, err
		}
		bs = append(bs, &books.Book{
			Id:     id,
			Title:  title,
			Author: author,
		})
	}
	if rows.Err() != nil {
		return bs, rows.Err()
	}
	return bs, nil
}

// Find selects one book by id from the database
func (r *Repo) Find(id uint64) (*books.Book, error) {
	var book *books.Book
	row := r.conn.QueryRow(context.Background(), "select title, author from books_old where id = $1", id)

	var title string
	var author string

	err := row.Scan(&title, &author)
	if err != nil {
		return book, err
	}
	book = &books.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}
	return book, err
}

// Create adds a new book to the databases
func (r *Repo) Create(book *books.Book) (*books.Book, error) {
	row := r.conn.QueryRow(context.Background(), "insert into books_old(title, author) values($1, $2) RETURNING id", book.Title, book.Author)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return book, err
	}
	book.Id = id
	return book, nil
}

// Update updates an existing record in the database
func (r *Repo) Update(book *books.Book) (*books.Book, error) {
	_, err := r.conn.Exec(context.Background(), "update books_old set title=$1, author=$2 where id=$3", book.Title, book.Author, book.Id)

	return book, err
}

// Delete deletes a record in the database
func (r *Repo) Delete(id uint64) error {
	_, err := r.conn.Exec(context.Background(), "delete from books_old where id=$1", id)

	return err
}
