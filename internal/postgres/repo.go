package postgres

import (
	"context"

	"github.com/benschw/books-poc/internal"
	"github.com/benschw/books-poc/models"
	"github.com/jackc/pgx/v4"
)

// Ensure Repo implements internal.BooksRepo.
var _ internal.Repo = &Repo{}

// Repo manages postgres database access for books
type Repo struct {
	conn *pgx.Conn
}

// NewRepo creates a new postgres repo
func NewRepo(dbStr string) (*Repo, error) {

	conn, err := pgx.Connect(context.Background(), dbStr)

	return &Repo{conn: conn}, err
}

// FindAll returns all books from the database
func (r *Repo) FindAll() ([]models.Book, error) {
	var bs []models.Book

	rows, _ := r.conn.Query(context.Background(), "select * from books order by id")

	for rows.Next() {
		var id uint64
		var title string
		var author string
		err := rows.Scan(&id, &title, &author)
		if err != nil {
			return bs, err
		}
		bs = append(bs, models.Book{
			ID:     id,
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
func (r *Repo) Find(id uint64) (models.Book, error) {
	var book models.Book
	row := r.conn.QueryRow(context.Background(), "select title, author from books where id = $1", id)

	var title string
	var author string

	err := row.Scan(&title, &author)
	if err != nil {
		return book, err
	}
	book = models.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
	return book, err
}

// Create adds a new book to the databases
func (r *Repo) Create(book models.Book) (models.Book, error) {
	row := r.conn.QueryRow(context.Background(), "insert into books(title, author) values($1, $2) RETURNING id", book.Title, book.Author)

	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return book, err
	}
	book.ID = id
	return book, nil
}

// Update updates an existing record in the database
func (r *Repo) Update(book models.Book) (models.Book, error) {
	_, err := r.conn.Exec(context.Background(), "update books set title=$1, author=$2 where id=$3", book.Title, book.Author, book.ID)

	return book, err
}

// Delete deletes a record in the database
func (r *Repo) Delete(id uint64) error {
	_, err := r.conn.Exec(context.Background(), "delete from books where id=$1", id)

	return err
}
