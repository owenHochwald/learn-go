package cli

import (
	"bookstore-cli/models"
	"bookstore-cli/services"
	"bookstore-cli/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	if len(os.Args) < 2 {
		fmt.Println("Expected at least one argument")
		os.Exit(1)
	}

	books := utils.LoadBooks()

	switch os.Args[1] {
	case "get":
		handleGet(books)
	case "update":
		books = handleUpdate(books)
	case "add":
		books = handleAdd(books)
	case "delete":
		books = handleDelete(books)

	default:
		fmt.Println("Unknown command")
	}
	utils.WriteBooks(books)

}

func handleGet(books []models.Book) {
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCommand.Bool("all", false, "Get all books")
	getId := getCommand.String("id", "", "Get book by id")

	getCommand.Parse(os.Args[2:])

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

}

func handleUpdate(books []models.Book) []models.Book {
	updateCommand := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCommand.String("id", "", "Update book by id")
	updateTitle := updateCommand.String("title", "", "Update book by id")
	updateAuthor := updateCommand.String("author", "", "Update book by id")
	updatePrice := updateCommand.String("price", "", "Update book by id")
	updateImageUrl := updateCommand.String("image_url", "", "Update book by id")

	updateCommand.Parse(os.Args[2:])
	books = services.UpdateBook(*updateId, *updateTitle, *updateAuthor, *updatePrice, *updateImageUrl, books)
	return books
}

func handleAdd(books []models.Book) []models.Book {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCommand.String("title", "", "add book by id")
	addAuthor := addCommand.String("author", "", "add book by id")
	addPrice := addCommand.String("price", "", "add book by id")
	addImageUrl := addCommand.String("image_url", "", "add book by id")

	addCommand.Parse(os.Args[2:])
	books = services.AddBook(*addTitle, *addAuthor, *addPrice, *addImageUrl, books)
	return books
}

func handleDelete(books []models.Book) []models.Book {
	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCommand.String("id", "", "Get book by id")

	deleteCommand.Parse(os.Args[2:])
	books = services.DeleteBook(*deleteId, books)
	return books
}
