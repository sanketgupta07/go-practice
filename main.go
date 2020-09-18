package main

import (
	closuers "github.com/sanketgupta07/go-practice/closures"
	"github.com/sanketgupta07/go-practice/recursion"
	multiplevalue "github.com/sanketgupta07/go-practice/return-multiple-value"
	variadic "github.com/sanketgupta07/go-practice/variadic_function"

	"log"
)

//Functions are First-class citizen in Go
var l = log.Println

func main() {

	l("Main function")
	l(multiplevalue.MultipleValue())
	l("Variadic- Sum: ", variadic.VariadicFunc(2, 5))
	l("Closures- CallClosures(10)", closuers.CallClosures(10))
	l("Recursion 1- Factorial (7): ", recursion.Factorial(7))
	l("Recursion 2- Fibonacci (5): ", recursion.CallFibonacci(5))
}
