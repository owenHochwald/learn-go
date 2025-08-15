package main

import (
	"fmt"
	"os"
)


func main () {

	// check file for errors
	

	args := os.Args[1:]

	if len(args) > 0 {
		curr := args[0]
		if curr == "add" {

		}
		fmt.Println(args[0])
	}

	fmt.Println(args)

}