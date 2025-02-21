package core

import (
	"math"
)

type HalfSpace int

const (
	HalfSpace_ON_PLANE HalfSpace = iota
	HalfSpace_BEHIND
	HalfSpace_FRONT
)

type Plane struct {
	Normal Vector3 // Normalized vector perpendicular to the plane
	Dist   float64 // Distance from a point to the plane.
}

// Constructs a "plane" based on 3 vectors
func BuildPlane(A Vector3, B Vector3, C Vector3) Plane {
	planeNormal := B.Sub(A).Cross(C.Sub(A)) // NORMAL = AB X AC
	planeNormal.Normalize()
	distFromPlane := -planeNormal.Dot(A)
	p := Plane{Normal: planeNormal, Dist: distFromPlane}
	return p
}

func (p Plane) RayIntersection(Start Vector3, Dir Vector3, time *float64, hitPt *Vector3) bool {
	denom := p.Normal.Dot(Dir)

	//If denom > 0 then the normal of the plane is pointing away from the ray
	//If denom==0 RAY AND PLANE ARE PARALLEL
	if math.Abs(denom) <= 1e-4 {
		return false
	}
	dist := -(p.Normal.Dot(Start) + p.Dist)
	tVal := dist / denom // CALC PARAMETRIC T-VAL
	time = &tVal
	if tVal > 0 { // IF INTERSECTION NOT "BEHIND" RAY
		hitPoint := Start.Add(Dir).Mul(tVal) // CALC HIT POINT
		hitPt = &hitPoint
		return true // INTERSECTION FOUND!
	}
	return false
}

func (p Plane) DistanceTo(point Vector3) float64 {
	return point.Dot(p.Normal) + p.Dist
}

func (p Plane) ClassifyPoint(point Vector3) HalfSpace {
	d := p.DistanceTo(point)
	if d == 0 {
		return HalfSpace_ON_PLANE
	}
	if d > 0 {
		return HalfSpace_FRONT
	}
	return HalfSpace_BEHIND
}
