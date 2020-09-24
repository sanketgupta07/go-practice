package main

import (
	"os"

	closuers "github.com/sanketgupta07/go-practice/closures"
	tree "github.com/sanketgupta07/go-practice/datastructure/tree"
	err "github.com/sanketgupta07/go-practice/err"
	recursion "github.com/sanketgupta07/go-practice/recursion"
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
	//Error call
	i, er1 := err.ErrorTest(-8)
	l("Error call :", i)
	if er1 != nil {
		l(er1.Error())
	} else {
		l(i)
	}

	//CustomTypeError call
	i, er2 := err.CustomErrExample(-7)
	l("Error call :", i)
	if er2 != nil {
		l(er2.Error())
	} else {
		l(i)
	}

	//Use data from error type
	if ae, ok := er2.(*err.CustomTypeErr); ok {
		l(ae.Arg)
		l(ae.Prob)
	}

	//Insert BinaryTree
	t := &tree.BinaryTree{}
	t.Insert(10).Insert(20).Insert(5).Insert(4).Insert(25).Insert(7).Insert(15).Insert(1)
	tree.PrintTree(os.Stdout, t.Root, 0, 'M')

}
