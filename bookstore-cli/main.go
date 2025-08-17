package main

import (
	"bookstore-cli/models"
	"bookstore-cli/services"
	"encoding/json"
	"fmt"
	"os"
)

const filePath = "books.json"

func main() {

	// check file for errors
	books := loadBooks()

	fmt.Println(books)

	services.GetAllBooks(books)

	services.GetBook(1, books)

	args := os.Args[1:]

	if len(args) > 0 {
		curr := args[0]
		if curr == "add" {

		}
		fmt.Println(args[0])
	}

	fmt.Println(args)

}

func loadBooks() []models.Book {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// create the file
		err := os.WriteFile(filePath, []byte("[]"), 0644)
		if err != nil {
			panic(err)
		}
		return []models.Book{}
	}

	// open file and read contents to memory in data
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// load books from file into books var
	var books []models.Book
	if err := json.Unmarshal(data, &books); err != nil {
		panic(err)
	}

	return books
}

func writeBooks(books []models.Book) {
	// get data and marshal into json
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		panic("Error marshalling books" + err.Error())
	}
	// write into books.json
	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		panic(err)
	}

	fmt.Println(("Saved books to books.json"))
}
