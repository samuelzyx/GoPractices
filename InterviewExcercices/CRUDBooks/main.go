package main

import (
	"CRUDBooks/controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/books", controller.GetBooks)
	http.HandleFunc("/books/get/:id", controller.GetBookByID)
	http.HandleFunc("/books/create", controller.CreateBook)
	http.HandleFunc("/books/update/:id", controller.UpdateBook)
	http.HandleFunc("/books/delete/:id", controller.DeleteBook)

	http.ListenAndServe(":8080", nil)

	fmt.Printf("Server running in PORT:%d\n", 8080)
}
