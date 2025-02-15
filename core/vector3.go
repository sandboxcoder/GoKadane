package core

import (
	"math"
)

type Vector3 struct {
	X, Y, Z float64 // Capitalized fields so they are exported outside of the package
}

// Returns a scalar (dot product between two vectors)
func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Returns the magnitude of a vector
func (v Vector3) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}
