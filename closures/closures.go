package closures

import (
	"log"
)

func closures(i int) func() int {
	log.Println("From closures initial value:", i)
	return func() int {
		i++
		return i
	}
	// log.Println("not reachable")
	// return closures(5)
	// uncomment above statements and run go vet ./... to see the out put for not reachable code
	// sanket@Khachedu MINGW64 ~/go/src/github.com/sanketgupta07/go-practice (master)
	// $ go vet ./...
	// # github.com/sanketgupta07/go-practice/closures
	// closures\closures.go:13:2: unreachable code

}

func CallClosures(i int) int {
	count := closures(i)
	log.Println("From CallClosures", count())
	log.Println("From CallClosures", count())
	log.Println("From CallClosures", count())
	log.Println("From CallClosures", count())
	return count()
}
