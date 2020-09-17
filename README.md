# go-practice
go tryouts for learning

Ref: https://gobyexample.com

Added - Makefile

Added - unit testing
    To run all test cases as test suite `go test -v`
    To run one test case  `go test -v -run <test-function-name>`
    To run all the testcase files recursively `go test ./...`

1. Multiple value return: 
Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

2. Variadic function: 
Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

3. Closures: Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
-- added go vet test for unreachable code in closures.go

4. Recursion: Go support recusrion.

    *   Factorial
    *   Fibonacci

