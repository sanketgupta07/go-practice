package main

import (
	"fmt"
	"time"
)

func routine(str string) {
	fmt.Println(str)
}

func main() {
	go routine("Hello from thread")
	go func(str string) {
		fmt.Println(str)
	}("Hello from anonymous function thread")
	routine("normal function call")

	time.Sleep(time.Second)
	fmt.Println("All function completed.")
}
