package main

import "fmt"

func main() {
	fmt.Println("Mad Lib 1.0")
	var noun, verb, adjective, adverb string

	fmt.Print("Enter a noun: ")
	fmt.Scan(&noun)

	fmt.Print("Now enter a verb: ")
	fmt.Scan(&verb)

	fmt.Print("Enter an adjective: ")
	fmt.Scan(&adjective)

	fmt.Print("Finally enter an adverb: ")
	fmt.Scan(&adverb)

	fmt.Println("Can you", verb, "your", adjective, noun, adverb)
}
