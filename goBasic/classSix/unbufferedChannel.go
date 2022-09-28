package main

import (
	"fmt"
	"time"
)

type Letter struct {
	Message string
}

func main() {
	// buffered channel -> indirect communication [mailbox]
	mailbox := make(chan Letter, 5)

	// process
	// -> postman sends letter
	go SendLetter(mailbox)

	// -> owner waits for it
	for {
		letter, ok := <-mailbox
		if !ok {
			break
		}

		fmt.Println("[Owner] Letter received")
		time.Sleep(time.Second * 5)
		fmt.Printf("[Owner] I've finished reading the letter and it says...`%s`", letter.Message)

	}

	fmt.Println("[Owner] No more letters in the box")
}

// routines
func SendLetter(out chan<- Letter) {
	for i:=0; i<10; i++ {
		time.Sleep(time.Second * 2)
		out <- Letter{fmt.Sprintf("New Letter! %d", i)}
		
		fmt.Println("[Postman] Letter sent")
	}
	close(out)
	fmt.Println("[Postman] No more letters available")
}