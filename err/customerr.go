package err

import (
	"fmt"
	"log"
)

/*
It’s possible to use custom types as errors by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
*/

type customTypeErr struct {
	arg  int
	prob string
}

//Error description for type customTypeErr
func (x *customTypeErr) Error() string {
	return fmt.Sprintf("%d - %s", x.arg, x.prob)
}

// CustomErrExample for custom error type
func CustomErrExample(i int) (int, error) {
	if i < 0 {
		return i, &customTypeErr{i, "I do not accept Negative value. "}
	}
	log.Println("I am happy with a positive or 0 value.", i)
	return i + 3, nil
}
