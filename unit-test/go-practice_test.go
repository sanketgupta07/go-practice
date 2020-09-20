package test

import (
	"testing"

	err "github.com/sanketgupta07/go-practice/err"

	closures "github.com/sanketgupta07/go-practice/closures"
	recursion "github.com/sanketgupta07/go-practice/recursion"

	variadic "github.com/sanketgupta07/go-practice/variadic_function"

	multipleValue "github.com/sanketgupta07/go-practice/return-multiple-value"
)

func TestMultipleValue(t *testing.T) {
	var expected1, expected2 int64
	expected1 = 33
	expected2 = 55
	actual1, actual2 := multipleValue.MultipleValue()
	if expected1 != actual1 || expected2 != actual2 {
		t.Errorf("return value is not equal to actual.")
	}

}

func TestVariadicFunc(t *testing.T) {
	expected := 10
	actual := variadic.VariadicFunc(7, 3)
	if expected != actual {
		t.Errorf("return value is not equal to actual.")
	}
}

func TestCallClosuers(t *testing.T) {
	expected := 8
	actual := closures.CallClosures(3)
	if expected != actual {
		t.Errorf("return value is not equal to actual.")
	}
}

func TestFactorial(t *testing.T) {
	expected := 120
	actual := recursion.Factorial(5)
	if expected != actual {
		t.Errorf("return value is not equal to actual.")
	}
}

func TestFibonacci(t *testing.T) {
	n := 5
	expected := []int{}
	expected = append(expected, 0, 1, 1, 2, 3)
	actual := recursion.CallFibonacci(n)
	for i := 0; i < n; i++ {
		if expected[i] != actual[i] {
			t.Errorf("return value is not equal to actual.")
		}

	}
}

func TestError(t *testing.T) {
	i1, i2 := 5, -3
	expected1 := 8
	actual1, er := err.ErrorTest(i1)

	expected2 := -3
	actual2, er := err.ErrorTest(i2)

	if expected1 != actual1 && er == nil {
		t.Errorf("return value is not equal to actual.")
	}

	if expected2 != actual2 && er != nil {
		t.Errorf("return value is not equal to actual.")
	}
}
