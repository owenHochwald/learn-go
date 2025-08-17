package main

import (
	"encoding/json"
	"fmt"
	"os"
	"bookstore-cli/models"
)

const filePath = "books.json"

func main() {

	// check file for errors
	books := loadBooks()

	fmt.Print(books)

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