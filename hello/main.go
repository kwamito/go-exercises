package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	if name == "kwame" {
		fmt.Println("Hello boss", name)
	} else if name == "nana" {
		fmt.Println("The main man", name, "is here")
	} else {
		fmt.Println("Hello", name)
	}
}
