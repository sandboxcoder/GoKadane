package core

import (
	"errors"
	"fmt"
	"math"
)

type BoundaryType int

const (
	Bounds_SPHERE BoundaryType = iota
	Bounds_BOX
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
		case Bounds_BOX:
			ret = a.Overlaps(b.Box)
		case Bounds_SPHERE:
			ret = a.sphereOverlaps(b)
		}
	} else {
		var err error
		ret, err = isPrimitivesCollidingSphereVsBox(a, b)
		if err != nil {
			fmt.Println("Error:", err)
			ret = a.Overlaps(b.Box)
		}
	}
	return ret
}

func (c CollisionPrimitive) GetSphereOrBox(sphere *Sphere, bbox *Box) error {
	var err error = nil
	if c.BoundaryType == Bounds_BOX {
		*bbox = c.Box
	} else if c.BoundaryType == Bounds_SPHERE {
		*sphere = c.GetSphere()
	} else {
		err = errors.New("CollisionPrimitive is not using a handled type")
	}
	return err
}

// Checks Box vs Sphere
func isPrimitivesCollidingSphereVsBox(a CollisionPrimitive, b CollisionPrimitive) (bool, error) {
	var err error = nil
	var ret bool = false
	if a.BoundaryType == b.BoundaryType {
		err = errors.New("CollisionPrimitive A & B should be using different types")
	} else {
		var bbox Box
		var sphere Sphere
		err = a.GetSphereOrBox(&sphere, &bbox)
		if err != nil {
			err = b.GetSphereOrBox(&sphere, &bbox)
		}
		ret = IsBoxCollidingWithSphere(bbox, sphere)
	}

	return ret, err
}

func IsBoxCollidingWithSphere(bbox Box, sphere Sphere) bool {
	closestPoint := Vector3{
		math.Max(bbox.Center.X-bbox.Extent.X, math.Min(sphere.Center.X, bbox.Center.X+bbox.Extent.X)),
		math.Max(bbox.Center.Y-bbox.Extent.Y, math.Min(sphere.Center.Y, bbox.Center.Y+bbox.Extent.Y)),
		math.Max(bbox.Center.Z-bbox.Extent.Z, math.Min(sphere.Center.Z, bbox.Center.Z+bbox.Extent.Z)),
	}
	distanceVec := sphere.Center.Sub(closestPoint)
	return distanceVec.Mag() <= sphere.Radius
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
