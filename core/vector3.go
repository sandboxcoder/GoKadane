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

func (v Vector3) LengthSquared() float64 {
	return (v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalizes the non-zero 3d vector
func (v *Vector3) Normalize() {
	magnitude := v.Mag()
	if magnitude == 0 {
		return // Avoid division by zero
	}
	v.X /= magnitude
	v.Y /= magnitude
	v.Z /= magnitude
}

func (a Vector3) Cross(b Vector3) Vector3 {
	return Vector3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{X: v1.X + v2.X, Y: v1.Y + v2.Y, Z: v1.Z + v2.Z}
}

// Subtraction operator
func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{X: v1.X - v2.X, Y: v1.Y - v2.Y, Z: v1.Z - v2.Z}
}

// Multiplication operator
func (v1 Vector3) Mul(scalar float64) Vector3 {
	return Vector3{X: v1.X * scalar, Y: v1.Y * scalar, Z: v1.Z * scalar}
}
