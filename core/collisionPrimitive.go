package core

type BoundaryType int

const (
	Bounds_SPHERE BoundaryType = iota
	Bounds_BoundingBox
)

// Encapsulates the collision boundary for an object.
type CollisionPrimitive struct {
	Box
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
		// Converting a sphere into a box is a fairly decent approx so we'll use that method.
		// Converting a box -> sphere would return an even bigger volume causing more false positives.
		ret = a.boxOverlaps(b)
	}
	return ret
}

func (c CollisionPrimitive) GetBoundingBox() Box {
	b := Box{c.Center, c.Extent}
	return b
}

func (a CollisionPrimitive) boxOverlaps(b CollisionPrimitive) bool {
	box2 := Box{b.Center, b.Extent}
	return a.Overlaps(box2)
}

func (a CollisionPrimitive) sphereOverlaps(b CollisionPrimitive) bool {
	s1 := a.GetSphere()
	s2 := b.GetSphere()
	return IsColliding(s1, s2)
}
