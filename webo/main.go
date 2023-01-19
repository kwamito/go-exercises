package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var link string

	fmt.Println("Website Scanner 1.0")

	fmt.Print("Enter the link to the website to be scanned: ")
	fmt.Scanln(&link)

	if link == "" {
		fmt.Println("No website provided, scanning default website")
		link = "https://collider.com/best-movies-that-turn-10-in-2023/"
	}

	fmt.Println("Scanning.......")
	fmt.Println("Website", link)

	response, err := get(link)

	if err != nil {
		fmt.Println(err)
		return
	}
	fo, errs := os.Create("result.html")

	if errs != nil {
		fmt.Println(errs)
		return
	}
	fo.Write([]byte(response))

}

func get(link string) (string, error) {
	resp, err := http.Get(link)

	if err != nil {
		fmt.Println(err, "Error returned")
		return "", errors.New("Error")
	}

	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(content))
	return string(content), nil
}

// func post(link string){
// 	resp, err := http.Post(link,"json",)

// 	if err != nil {
// 		fmt.Println(err, "Error returned")
// 		return
// 	}

// 	content, err := ioutil.ReadAll(resp.Body)
// 	fmt.Println("Response: ", string(content))
// 	return
// }
