package core

import "math"

// Axis aligned bounding box
type Box struct {
	Center Vector3
	Extent Vector3
}

func (b Box) Overlaps(box Box) bool {
	dir := box.Center.Sub(b.Center) //vector from A to B
	return (math.Abs(dir.X) <= (b.Extent.X + box.Extent.X)) && (math.Abs(dir.Y) <= (b.Extent.Y + box.Extent.Y)) && (math.Abs(dir.Z) <= (b.Extent.Z + box.Extent.Z))
}

func (b Box) GetMin() Vector3 {
	min := b.Center.Sub(b.Extent)
	return min
}

func (b Box) GetMax() Vector3 {
	max := b.Center.Add(b.Extent)
	return max
}

func (b Box) Size() float64 {
	diagonal := b.Extent.Mag()
	return diagonal * 2
}

func (b Box) GetSphere() Sphere {
	var sphere Sphere

	//Converts AABB to sphere
	sphere.Center = b.Center
	sphere.Radius = b.Extent.Mag()
	return sphere
}
