package main

// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus


// ---------------------------------------------------------
// EXERCISE: Print names
//
//  Print your name and your best friend's name using
//  Println twice
//
// EXPECTED OUTPUT
//  YourName
//  YourBestFriendName
//
// BONUS
//  Use `go run` first.
//  And after that use `go build` and run your program.
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	fmt.Println("Owen Hochwald")
	fmt.Println("Faheem Ahmed")

	fmt.Printf("Hello I am printing this: %d\n", 9)
}