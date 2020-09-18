package variadic_function

import (
	"log"
)

//VariadicFunc: SUM return the sum of all input int
func VariadicFunc(nums ...int) int {
	log.Println("from VariadicFunc.")
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	return sum
}
