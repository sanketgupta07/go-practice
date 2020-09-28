package iface

// Measure Returns Area and Perimeter
func Measure(g Geometry) (float64, float64) {
	a := g.Area()
	p := g.Peri()
	return a, p
}
