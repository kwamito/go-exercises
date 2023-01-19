package main

import (
	"fmt"
)

func main() {
	fmt.Println("Simple Math 1.0")

	var firstNumber, lastNumber int

	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter last number: ")
	fmt.Scan(&lastNumber)

	calculate(firstNumber, lastNumber)

}

func calculate(firstNumber int, lastNumber int) {
	// var result float32
	// if firstNumber < 0 {
	// 	// fmt.Println("Number less than zero is not allowed")
	// 	return result, errors.New("numner less than zero is not allowed")
	// }
	if firstNumber < 0 {
		fmt.Println("numbers less than zero not allowed")
		return
	}
	if lastNumber < 0 {
		fmt.Println("numbers less than zero not allowed")
		return
	} else {
		fmt.Println("Addition -", firstNumber+lastNumber)
		fmt.Println("Subtraction -", firstNumber-lastNumber)
		fmt.Println("Multiplication -", firstNumber*lastNumber)
		fmt.Println("Division -", firstNumber/lastNumber)
	}

}
