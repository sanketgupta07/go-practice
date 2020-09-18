package multiplevalue

import "log"

//MultipleValue return two int64
func MultipleValue() (int64, int64) {
	log.Println("From multiplereturn function")
	return 33, 55
}
