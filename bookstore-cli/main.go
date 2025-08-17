package main

import (
	"bookstore-cli/services"
	"bookstore-cli/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCommand.Bool("all", false, "Get all books")
	getId := getCommand.String("id", "", "Get book by id")

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCommand.String("title", "", "add book by id")
	addAuthor := addCommand.String("author", "", "add book by id")
	addPrice := addCommand.String("price", "", "add book by id")
	addImageUrl := addCommand.String("image_url", "", "add book by id")

	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCommand.String("id", "", "Update book by id")
	updateTitle := updateCommand.String("title", "", "Update book by id")
	updateAuthor := updateCommand.String("author", "", "Update book by id")
	updatePrice := updateCommand.String("price", "", "Update book by id")
	updateImageUrl := updateCommand.String("image_url", "", "Update book by id")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCommand.String("id", "", "Get book by id")

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Expected at least one argument")
		os.Exit(1)
	}

	// check file for errors
	books := utils.LoadBooks()

	switch args[1] {
	case "get":
		getCommand.Parse(args[2:])

		switch {
		case *getAll:
			services.GetAllBooks(books)

		case *getId != "":
			id, err := strconv.Atoi(*getId)
			if err != nil {
				fmt.Println("Expected a number")
			}
			services.GetBook(id, books)

		default:
			fmt.Println("Usage: get --all OR get --id <id>")
		}
	case "add":
		addCommand.Parse(args[2:])
		books = services.AddBook(*addTitle, *addAuthor, *addPrice, *addImageUrl, books)
	case "update":
		updateCommand.Parse(args[2:])
		books = services.UpdateBook(*updateId, *updateTitle, *updateAuthor, *updatePrice, *updateImageUrl, books)
	case "delete":
		deleteCommand.Parse(args[2:])
		books = services.DeleteBook(*deleteId, books)

	default:
		fmt.Println("Unknown command")
	}

	utils.WriteBooks(books)

}
