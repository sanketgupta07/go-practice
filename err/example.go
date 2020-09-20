package err

import (
	"errors"
	"log"
)

func ErrorTest(i int) (int, error) {
	if i < 0 {
		return i, errors.New("I do not accept Negative value. ")
	}
	log.Println("I am happy with a positive or 0 value.", i)
	return i + 3, nil
}
