package main

import (
	"fmt"
	"time"
)

type Letter struct {
	Message string
}

func main() {
	// unbuffered channel -> direct communication
	meet := make(chan Letter)


	// process
	// -> postman sends letter
	go SendLetter(meet)

	// -> owner waits for it
	letter := <- meet
	fmt.Println("[Owner] Letter received")

	time.Sleep(time.Second * 5)
	fmt.Printf("[Owner] I've finished reading the letter and it says...`%s`", letter.Message)
}

// routines
func SendLetter(out chan<- Letter) {
	time.Sleep(time.Second * 5)

	out <- Letter{"New Letter!"}
	fmt.Println("[Postman] Letter sent")
	close(out)
}