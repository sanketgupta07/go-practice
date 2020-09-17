package recursion

import (
	"log"
)

func fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 || i == 2 {
		return 1
	}

	return fibonacci(i-1) + fibonacci(i-2)
}

//CallFibonacci
func CallFibonacci(n int) []int {
	log.Print("From CallFibonacci ")
	series := []int{}
	for i := 0; i < n; i++ {
		series = append(series, fibonacci(i))
	}
	return series
}
