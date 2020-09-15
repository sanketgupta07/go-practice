package test

import (
	"testing"

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
