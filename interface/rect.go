package iface

type Rect struct {
	Length float64
	Width  float64
}

func (r Rect) Area() float64 {
	return r.Length * r.Width
}

func (r Rect) Peri() float64 {
	return 2 * (r.Length + r.Width)
}
