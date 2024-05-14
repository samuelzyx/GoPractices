package repository

import (
	"CRUDBooks/model"
	"errors"
)

var books = make(map[int]model.Book)
var nextID = 1

func GetBooks() []model.Book {
	result := make([]model.Book, 0, len(books))
	for _, book := range books {
		result = append(result, book)
	}
	return result
}

func GetBookByID(id int) (model.Book, error) {
	book, ok := books[id]
	if !ok {
		return model.Book{}, errors.New("book not found")
	}
	return book, nil
}

func CreateBook(book model.Book) (int, error) {
	book.ID = nextID

	books[nextID] = book
	nextID++
	return book.ID, nil
}

func UpdateBook(id int, newBook model.Book) error {
	_, ok := books[id]
	if !ok {
		return errors.New("book not found")
	}
	books[id] = newBook
	return nil
}

func DeleteBook(id int) error {
	_, ok := books[id]
	if !ok {
		return errors.New("book not found")
	}
	delete(books, id)
	return nil
}
