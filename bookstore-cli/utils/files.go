package utils

import (
	"bookstore-cli/models"
	"encoding/json"
	"fmt"
	"os"
)

const filePath = "books.json"

func LoadBooks() []models.Book {
	createFile()

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var books []models.Book
	if err := json.Unmarshal(data, &books); err != nil {
		panic(err)
	}

	return books
}

func createFile() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.WriteFile(filePath, []byte("[]"), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func WriteBooks(books []models.Book) {
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
