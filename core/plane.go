package core

type HalfSpace int

const (
	HalfSpace_BEHIND HalfSpace = iota
	HalfSpace_ON_PLANE
	HalfSpace_FRONT
)

type Plane struct {
	Normal Vector3
	Dist   float64
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

func (p Plane) ClassifyPoint(pt Vector3) HalfSpace {
	return HalfSpace_BEHIND
}
