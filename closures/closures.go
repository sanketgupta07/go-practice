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
}

func CallClosures(i int) int {
	count := closures(i)
	log.Println("From CallClosures", count())
	log.Println("From CallClosures", count())
	log.Println("From CallClosures", count())
	log.Println("From CallClosures", count())
	return count()
}
