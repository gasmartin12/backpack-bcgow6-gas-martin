package main

import (
	"fmt"
	"os"
	"time"

	"context"
)

func Function(ctx context.Context, data int) {
	value := ctx.Value("Key")

	str, ok := value.(int)
	if !ok {
		fmt.Println("Not string")
		os.Exit(1)
	}

	fmt.Printf("%T", str)
}

func Function2(ctx context.Context, data int) {
	value := ctx.Value("Key")

	str, ok := value.(string)
	if !ok {
		fmt.Println("Not string")
		os.Exit(1)
	}

	fmt.Printf("%T", str)
	_, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

}

func main() {
	ctx := context.Background()

	ctx2 := context.WithValue(ctx, "Key", "Value")

	//Fuction(ctx2, 10)
	//fmt.Printf("%+v", ctx)
	//ctx3, _ := context.WithDeadline(ctx2, time.Now().Add(5*time.Second))
	//<-ctx3.Done()

	//fmt.Println(ctx3.Err().Error())

	//ctx4, cancel := context.WithTimeout(ctx2, 5*time.Second)

	Function2(ctx2, 10)

	<-ctx2.Done()
	fmt.Println(ctx2.Err().Error())

}
