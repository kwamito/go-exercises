package main

import (
	"fmt"
	"strings"

	"github.com/rivo/uniseg"
)

func main() {
	fmt.Println("Character counter 1.0")
	fmt.Print("Enter a word to count: ")
	var word string
	fmt.Scanln(&word)

	count := strings.Count(word, "")
	fmt.Println(count, "This is the count")
	fmt.Println(uniseg.GraphemeClusterCount(word), "real")
}
