package core

import (
	"fmt"
)

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

func (c CollisionPrimitive) GetSphere() Sphere {
	var s Sphere
	if c.BoundaryType == Bounds_BoundingBox {
		box := Box{c.Center, c.Extent}
		s = box.GetSphere()
	} else if c.BoundaryType == Bounds_SPHERE {
		s = Sphere{c.Center, c.Extent.X}
	} else {
		s = Sphere{c.Center, c.Extent.X}
		fmt.Printf("Warn: Unknown primitive type %+v", c.BoundaryType)
	}
	return s
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
