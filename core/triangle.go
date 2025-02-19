package core

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
