package main

import (
	"fmt"
)

func main() {

	word := "palabra"

	wordArray := []rune(word)
	countLetters := len(wordArray)

	fmt.Println("The word '", word, "'")
	fmt.Println("has ", countLetters, " characters")
	for i := 0; i < countLetters; i++ {
		fmt.Println(string(wordArray[i]))
	}

}
