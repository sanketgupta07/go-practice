package iface

import "fmt"

func Measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Peri())
}
