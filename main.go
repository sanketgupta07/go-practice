package main

import (
	closuers "github.com/sanketgupta07/go-practice/closures"
	multiplevalue "github.com/sanketgupta07/go-practice/return-multiple-value"
	variadic "github.com/sanketgupta07/go-practice/variadic_function"

	"log"
)

func main() {

	log.Println("Main function")
	log.Println(multiplevalue.MultipleValue())
	log.Println("Sum: ", variadic.VariadicFunc(2, 5))
	log.Println("Calling CallClosures", closuers.CallClosures(10))
}
