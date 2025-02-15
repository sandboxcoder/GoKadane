package core

type Vector3 struct {
	X, Y, Z float64 // Capitalized fields so they are exported outside of the package
}

func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}
