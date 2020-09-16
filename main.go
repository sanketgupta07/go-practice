package main

import (
<<<<<<< HEAD
	closuers "github.com/sanketgupta07/go-practice/closures"
=======
>>>>>>> cba3d14966d59ac5e65ad458deb1bfbcf2b3173f
	multiplevalue "github.com/sanketgupta07/go-practice/return-multiple-value"
	variadic "github.com/sanketgupta07/go-practice/variadic_function"

	"log"
)

func main() {

	log.Println("Main function")
	log.Println(multiplevalue.MultipleValue())
	log.Println("Sum: ", variadic.VariadicFunc(2, 5))
<<<<<<< HEAD
	log.Println("Calling CallClosures", closuers.CallClosures(10))
=======
>>>>>>> cba3d14966d59ac5e65ad458deb1bfbcf2b3173f
}
