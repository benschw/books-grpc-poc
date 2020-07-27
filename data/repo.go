package data

import (
	"context"

	"github.com/benschw/books-poc/api"
	"github.com/jackc/pgx/v4"
)

// BooksRepo manages database access for books
type BooksRepo struct {
	Conn *pgx.Conn
}

// FindAll returns all books from the database
func (r *BooksRepo) FindAll() ([]api.Book, error) {
	var books []api.Book

	rows, _ := r.Conn.Query(context.Background(), "select * from books order by id")

	for rows.Next() {
		var id uint64
		var title string
		var author string
		err := rows.Scan(&id, &title, &author)
		if err != nil {
			return books, err
		}
		books = append(books, api.Book{
			ID:     id,
			Title:  title,
			Author: author,
		})
	}
	if rows.Err() != nil {
		return books, rows.Err()
	}
	return books, nil
}

// Find selects one book by id from the database
func (r *BooksRepo) Find(id uint64) (api.Book, error) {
	var book api.Book
	row := r.Conn.QueryRow(context.Background(), "select title, author from books where id = $1", id)

	var title string
	var author string

	err := row.Scan(&title, &author)
	if err != nil {
		return book, err
	}
	book = api.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
	return book, err
}

// Create adds a new book to the databases
func (r *BooksRepo) Create(book api.Book) (api.Book, error) {
	row := r.Conn.QueryRow(context.Background(), "insert into books(title, author) values($1, $2) RETURNING id", book.Title, book.Author)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return book, err
	}
	book.ID = id
	return book, nil
}

// Update updates an existing record in the database
func (r *BooksRepo) Update(book api.Book) (api.Book, error) {
	_, err := r.Conn.Exec(context.Background(), "update books set title=$1, author=$2 where id=$3", book.Title, book.Author, book.ID)

	return book, err
}

// Delete deletes a record in the database
func (r *BooksRepo) Delete(id uint64) error {
	_, err := r.Conn.Exec(context.Background(), "delete from books where id=$1", id)

	return err
}
