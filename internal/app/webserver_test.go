package app

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/benschw/books-poc/internal/fakes"
	"github.com/benschw/books-poc/models"
	"github.com/stretchr/testify/assert"
)

var input []models.Book = []models.Book{
	models.Book{Title: "hello world", Author: "Ben"},
	models.Book{Title: "hello galaxy", Author: "Schwartz"},
}

var expected []models.Book = []models.Book{
	models.Book{ID: 1, Title: "hello world", Author: "Ben"},
	models.Book{ID: 2, Title: "hello galaxy", Author: "Schwartz"},
}

func fakeRequest(r *WebServer, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.Router().ServeHTTP(w, req)
	return w
}

func TestFindAllBooks(t *testing.T) {
	// given
	repo := fakes.NewRepo()
	repo.Create(input[0])
	repo.Create(input[1])

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "GET", "/book", nil)

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]models.Book
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	books, exists := response["data"]
	assert.Nil(t, err)
	assert.True(t, exists)

	assert.Equal(t, expected[0].ID, books[0].ID)
	assert.Equal(t, expected[0].Title, books[0].Title)
	assert.Equal(t, expected[0].Author, books[0].Author)

	assert.Equal(t, expected[1].ID, books[1].ID)
	assert.Equal(t, expected[1].Title, books[1].Title)
	assert.Equal(t, expected[1].Author, books[1].Author)
}

func TestCreateBook(t *testing.T) {
	// given
	repo := fakes.NewRepo()

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "POST", "/book", strings.NewReader("{\"title\": \"hello world\", \"author\": \"Ben\" }"))

	// then
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]models.Book
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	book, exists := response["data"]
	assert.Nil(t, err)
	assert.True(t, exists)

	assert.Equal(t, expected[0].ID, book.ID)
	assert.Equal(t, expected[0].Title, book.Title)
	assert.Equal(t, expected[0].Author, book.Author)

	assert.Equal(t, expected[0].ID, repo.Books[0].ID)
	assert.Equal(t, expected[0].Title, repo.Books[0].Title)
	assert.Equal(t, expected[0].Author, repo.Books[0].Author)
}

func TestFindBook(t *testing.T) {
	// given
	repo := fakes.NewRepo()
	repo.Create(input[0])

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "GET", "/book/1", nil)

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.Book
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	book, exists := response["data"]
	assert.Nil(t, err)
	assert.True(t, exists)

	assert.Equal(t, expected[0].ID, book.ID)
	assert.Equal(t, expected[0].Title, book.Title)
	assert.Equal(t, expected[0].Author, book.Author)
}

func TestFindBook_NotFound(t *testing.T) {
	// given
	repo := fakes.NewRepo()

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "GET", "/book/1", nil)

	// then
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateBook(t *testing.T) {
	// given
	repo := fakes.NewRepo()
	repo.Create(input[1])

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "PUT", "/book/1", strings.NewReader("{\"title\": \"hello world\", \"author\": \"Ben\" }"))

	// then
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.Book
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	book, exists := response["data"]
	assert.Nil(t, err)
	assert.True(t, exists)

	assert.Equal(t, expected[0].ID, book.ID)
	assert.Equal(t, expected[0].Title, book.Title)
	assert.Equal(t, expected[0].Author, book.Author)

	assert.Equal(t, expected[0].ID, repo.Books[0].ID)
	assert.Equal(t, expected[0].Title, repo.Books[0].Title)
	assert.Equal(t, expected[0].Author, repo.Books[0].Author)
}

func TestUpdateBook_NotFound(t *testing.T) {
	// given
	repo := fakes.NewRepo()

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "PUT", "/book/1", strings.NewReader("{\"title\": \"hello galaxy\", \"author\": \"Schwartz\" }"))

	// then
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteBook(t *testing.T) {
	// given
	repo := fakes.NewRepo()
	repo.Create(models.Book{Title: "hello world", Author: "Ben"})

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "DELETE", "/book/1", nil)

	// then
	assert.Equal(t, http.StatusNoContent, w.Code)

	assert.Equal(t, 0, len(repo.Books))
}

func TestDeleteBook_NotFound(t *testing.T) {
	// given
	repo := fakes.NewRepo()

	webApp := NewWebServer(repo)

	// when
	w := fakeRequest(webApp, "DELETE", "/book/1", nil)

	// then
	assert.Equal(t, http.StatusNotFound, w.Code)
}
