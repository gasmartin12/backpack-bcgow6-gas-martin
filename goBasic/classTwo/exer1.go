package main

import (
	"fmt"
)

func main() {

	var word string
	fmt.Print("Insert a word : ")
	fmt.Scanf("%s", &word)

	wordArray := []rune(word)
	letters := len(wordArray)

	fmt.Println("The word '", word, "'")
	fmt.Println("has ", letters, " characters")
	for i := 0; i < letters; i++ {
		fmt.Println(string(wordArray[i]))
	}

}
