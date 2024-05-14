package service

import (
	"CRUDBooks/model"
	"CRUDBooks/repository"
	"testing"
)

func TestGetBookByID(t *testing.T) {
	testBook := model.Book{
		ID:     1,
		Title:  "Test Book",
		Author: "Test Author",
	}

	repository.CreateBook(testBook)

	book, err := GetBookByID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if book != testBook {
		t.Errorf("expected book: %v, got: %v", testBook, book)
	}

	_, err = GetBookByID(2)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if err.Error() != "book not found" {
		t.Errorf("expected error: 'book not found', got: %v", err)
	}
}
