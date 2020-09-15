package main

import (
	multiplevalue "github.com/sanketgupta07/go-practice/return-multiple-value"
	variadic "github.com/sanketgupta07/go-practice/variadic_function"

	"log"
)

func main() {

	log.Println("Main function")
	log.Println(multiplevalue.MultipleValue())
	log.Println("Sum: ", variadic.VariadicFunc(2, 5))
}
