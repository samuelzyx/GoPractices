package service

import (
	"CRUDBooks/model"
	"CRUDBooks/repository"
	"errors"
)

func GetAllBooks() []model.Book {
	return repository.GetBooks()
}

func GetBookByID(id int) (model.Book, error) {
	return repository.GetBookByID(id)
}

func CreateBook(book model.Book) (int, error) {
	if book.Title == "" || book.Author == "" {
		return 0, errors.New("title and author are required")
	}

	return repository.CreateBook(book)
}

func UpdateBook(id int, newBook model.Book) error {
	return repository.UpdateBook(id, newBook)
}

func DeleteBook(id int) error {
	return repository.DeleteBook(id)
}
