package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader(
		"Hello world",
	)
	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	io.WriteString(
		os.Stdout,
		"Hello world! :)")

}
