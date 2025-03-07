package core

type BoundaryType int

const (
	Bounds_SPHERE BoundaryType = iota
	Bounds_BoundingBox
)

// Encapsulates the collision boundary for an object.
type CollisionPrimitive struct {
	Box // This variable might be reused to store a Sphere. Query BoundaryType value to determine which shape it is
	BoundaryType
}

func (a CollisionPrimitive) IsColliding(b CollisionPrimitive) bool {
	var ret bool
	if a.BoundaryType == b.BoundaryType {
		// This is the ideal scenario, where the collision primitives are the same type and no conversion
		// is needed.
		switch a.BoundaryType {
		case Bounds_BoundingBox:
			ret = a.Overlaps(b.Box)
		case Bounds_SPHERE:
			ret = a.sphereOverlaps(b)
		}
	} else {
		// Converting a sphere into a box is not an ideal approx but we'll have to use that method.
		// Converting a box -> sphere returns an even bigger volume which can trigger false positives.
		// It's best not to compare different shapes since we don't have an actual Sphere vs Box algorithm.
		ret = a.Overlaps(b.Box)
	}
	return ret
}

func (c CollisionPrimitive) GetSphere() Sphere {
	var s Sphere
	if c.BoundaryType == Bounds_SPHERE {
		// In this case, the primitive contains a sphere. Do not invoke the Box.GetSphere()
		// function which will just return an "enlarged" sphere. Instead, build the sphere manually
		// because the X-axis stores the raw Sphere radius.
		s = Sphere{c.Center, c.Extent.X}
	} else {
		s = c.Box.GetSphere()
	}
	return s
}

func (a CollisionPrimitive) sphereOverlaps(b CollisionPrimitive) bool {
	s1 := a.GetSphere()
	s2 := b.GetSphere()
	return IsColliding(s1, s2)
}
