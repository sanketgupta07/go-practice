package test

import (
	"fmt"
	"testing"

	iface "github.com/sanketgupta07/go-practice/interface"

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

func TestInterface(t *testing.T) {
	var ea1, ep1, ea2, ep2 float64
	ea1 = 10
	ep1 = 14

	ea2 = 314.1592653589793
	ep2 = 62.83185307179586
	rect := iface.Rect{Length: 2, Width: 5}
	circle := iface.Circle{Radius: 10}

	a1, p1 := iface.Measure(rect)
	a2, p2 := iface.Measure(circle)

	if ea1 != a1 && ep1 != p1 {
		fmt.Println(a1, p1)
		t.Errorf("return value is not equal to actual.")
	}

	if ea2 != a2 && ep2 != p2 {
		t.Errorf("return value is not equal to actual.")
	}

}
