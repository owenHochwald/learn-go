package main

import (
	"bookstore-cli/services"
	"bookstore-cli/utils"
	"fmt"
	"os"
)

func main() {

	// check file for errors
	books := utils.LoadBooks()

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

	utils.WriteBooks(books)

}
