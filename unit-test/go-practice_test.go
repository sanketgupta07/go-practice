package test

import (
	"testing"

<<<<<<< HEAD
	closures "github.com/sanketgupta07/go-practice/closures"

=======
>>>>>>> cba3d14966d59ac5e65ad458deb1bfbcf2b3173f
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
<<<<<<< HEAD

func TestCallClosuers(t *testing.T) {
	expected := 8
	actual := closures.CallClosures(3)
	if expected != actual {
		t.Errorf("return value is not equal to actual.")
	}
}
=======
>>>>>>> cba3d14966d59ac5e65ad458deb1bfbcf2b3173f
