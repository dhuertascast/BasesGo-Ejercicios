package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter th word to evaluate")
	var word string
	fmt.Scanln(&word)
	numberOfLetters := len(word)
	fmt.Printf("The number of letters of the word  '%s' is %d. \n", word, numberOfLetters)
	for _, char := range word {
		fmt.Println(string(char))
	}
}
