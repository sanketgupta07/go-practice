package recursion

func Factorial(i int) int {
	if i == 0 {
		return 1
	}

	return i * Factorial(i-1)
}
