package main

import "fmt"

func main() {
	fmt.Println("Quotes")
	fmt.Print("Enter quote and author: ")

	var quote, author string

	fmt.Scan(&quote,&author)

	// fmt.Print("Enter author: ")
	// fmt.Scanln(&author)

	fmt.Println(author, "says", quote)
}