package main

import (
	"fmt"
	// "reflect"
	"time"
)

func main() {
	fmt.Println("Retirement Calculator 1.0")

	var age, ageToRetire int

	currentYear := time.Now().Year()

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter the age you'll like to retire: ")
	fmt.Scan(&ageToRetire)

	// if reflect.TypeOf(age) ==
	// fmt.Println(reflect.TypeOf(age).Kind(), "This is the type of the var")

	ageDiff := ageToRetire - age
	yearToRetire := ageDiff + currentYear

	if ageDiff <= 0 {
		fmt.Println("You can already retire")
		return
	}

	// fmt.Println("Your age is ", ageDiff, currentYear, yearToRetire)
	fmt.Println("You have", ageDiff, "years left until you can retire.")
	fmt.Println("It's", currentYear, ",so you can retire in", yearToRetire, ".")
}
