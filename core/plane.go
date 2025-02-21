package core

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

func BuildPlane(A Vector3, B Vector3, C Vector3) Plane {
	n := B.Sub(A).Cross(C.Sub(A)) // NORMAL = AB X AC
	n.Normalize()
	distFromPlane := -n.Dot(A)
	p := Plane{Normal: n, Dist: distFromPlane}
	return p
}

func (p Plane) RayIntersection(Start Vector3, Dir Vector3, t *float64, hitPt *Vector3) bool {
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
