package library

import (
	"fmt"
	"goapp-giahuy/learninggo/day16/models"
)

type Library struct {
	Books map[string]models.Book
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]models.Book),
	}
}

func (lib *Library) AddBookStore(id, title, author string) error {
	if _, exists := lib.Books[id]; exists {
		return fmt.Errorf("sach voi Id %s da ton tai", id)
	}

	lib.Books[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}

	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.Books))

	for _, book := range lib.Books {
		books = append(books, book)
	}

	return books
}
