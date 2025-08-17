package main

import (
	"bookstore-cli/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const filePath = "books.json"

func main() {

	// check file for errors
	books := loadBooks()

	fmt.Println(books)

	getAllBooks(books)

	getBook(1, books)

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

func getAllBooks(books []models.Book) {
	fmt.Println("Id\t\tTitle\t\tAuthor\t\tImage URL")

	// print all books to console
	for _, book := range books {
		printBook(book)
	}
	fmt.Println()

}

func getBook(id int, books []models.Book) {
	for _, book := range books {
		bookId, err := strconv.Atoi(book.Id)
		if err != nil {
			panic(err)
		}

		if bookId == id {
			fmt.Println("Id\t\tTitle\t\tAuthor\t\tImage URL")
			printBook(book)
			fmt.Println()
		}
	}
}

func printBook(book models.Book) {
	fmt.Printf("%s\t%s\t%s\t%s\n", book.Id, book.Title, book.Author, book.Image_url)
}
