package core

type Triangle struct {
	base, height float64
}

func Area(t Triangle) float64 {
	return 0.5 * t.base * t.height
}

func NewTriangle(base, height float64) Triangle {
	t := Triangle{base, height}
	return t
}
