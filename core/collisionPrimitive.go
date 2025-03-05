package core

type BoundaryType int

const (
	Bounds_SPHERE BoundaryType = iota
	Bounds_BoundingBox
)

// Encapsulates the collision boundary for an object.
type CollisionPrimitive struct {
	Center Vector3
	Extent Vector3
	BoundaryType
}

func (a CollisionPrimitive) IsColliding(b CollisionPrimitive) bool {
	var ret bool
	if a.BoundaryType == b.BoundaryType {
		// This is the ideal scenario, where the collision primitives are the same type and no conversion
		// is needed.
		switch a.BoundaryType {
		case Bounds_BoundingBox:
			ret = a.boxOverlaps(b)
		case Bounds_SPHERE:
			ret = a.sphereOverlaps(b)
		}
	} else {
		ret = a.sphereOverlaps(b)
	}
	return ret
}

func (a CollisionPrimitive) boxOverlaps(b CollisionPrimitive) bool {
	box1 := Box{a.Center, a.Extent}
	box2 := Box{b.Center, b.Extent}
	return box1.Overlaps(box2)
}

func (a CollisionPrimitive) sphereOverlaps(b CollisionPrimitive) bool {
	s1 := Sphere{a.Center, a.Extent.X}
	s2 := Sphere{b.Center, b.Extent.X}
	return IsColliding(s1, s2)
}
