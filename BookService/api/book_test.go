package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Spring Boot in Action", Author: "Craig Walls", ISBN: "0345391802"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Spring Boot in Action","author":"Craig Walls","isbn":"0345391802"}`, string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Spring Boot in Action","author":"Craig Walls","isbn":"0345391802"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "Spring Boot in Action", Author: "Craig Walls", ISBN: "0345391802"}, book, "Book JSON unmarshalling wrong.")
}

func TestAllBooks(t *testing.T) {
	books := ListBooks()
	assert.Len(t, books, 4, "Wrong number of books.")
}

func TestCreateNewBook(t *testing.T) {
	book := Book{Title: "Test", Author: "Me", ISBN: "1234567890"}
	isbn, created := CreateBook(book)
	assert.True(t, created, "Book was not created.")
	assert.Equal(t, "1234567890", isbn, "Wrong ISBN.")
}

func TestDoNotCreateExistingBook(t *testing.T) {
	book := Book{ISBN: "0345391802"}
	_, created := CreateBook(book)
	assert.False(t, created, "Book was created.")
}

func TestUpdateExistingBook(t *testing.T) {
	book := Book{Title: "Test Update", Author: "Me Again", ISBN: "1234567890", Description: "Dummy Text"}
	updated := UpdateBook("1234567890", book)
	assert.True(t, updated, "Book not updated.")

	book, _ = GetBook("1234567890")
	assert.Equal(t, "Test Update", book.Title, "Title not updated.")
	assert.Equal(t, "Me Again", book.Author, "Author not updated.")
	assert.Equal(t, "Dummy Text", book.Description, "Description not updated.")
}

func TestDeleteBook(t *testing.T) {
	DeleteBook("1234567890")
	assert.Len(t, ListBooks(), 4, "Wrong number of books after delete.")
}