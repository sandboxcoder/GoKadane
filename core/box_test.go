package core

import (
	"testing"
)

func Test_BoxOverlaps(t *testing.T) {
	center := Vector3{0, 0, 0}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	detectedCollision := b1.Overlaps(b1)
	if detectedCollision != true {
		t.Errorf("overlaps() != expected, returned %+v instead of true", detectedCollision)
	}
}

func Test_BoxMax(t *testing.T) {
	center := Vector3{1, 1, 1}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	max := b1.GetMax()
	expected := Vector3{6, 6, 6}
	if max != expected {
		t.Errorf("GetMax() != expected, returned %+v instead of expected value %f", max, expected)
	}
}

func Test_BoxMin(t *testing.T) {
	center := Vector3{1, 1, 1}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	result := b1.GetMin()
	expected := Vector3{-4, -4, -4}
	if result != expected {
		t.Errorf("result != expected, returned %+v instead of expected value %f", result, expected)
	}
}

func Test_BoxExtent(t *testing.T) {
	center := Vector3{0, 0, 0}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	result := b1.GetSize()
	max := b1.GetMax()
	min := b1.GetMin()
	expected := (max.Sub(min)).Mag()
	if result != expected {
		t.Errorf("Size() != expected, returned %+v instead of expected value %f", result, expected)
	}
}

func TestGetSphere_Center(t *testing.T) {
	center := Vector3{0, 0, 0}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	sphere := b1.GetSphere()
	result := sphere.Center == b1.Center
	expected := true
	if result != expected {
		t.Errorf("GetSphere() != expected, returned %+v instead of expected value %+v", result, expected)
	}
}

func TestGetSphere_Radius(t *testing.T) {
	center := Vector3{0, 0, 0}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	sphere := b1.GetSphere()
	// Verify the sphere encloses the box
	result := sphere.Radius > b1.Extent.X && sphere.Radius > b1.Extent.Y && sphere.Radius > b1.Extent.Z
	expected := true
	if result != expected {
		t.Errorf("GetSphere() != expected, returned %+v instead of expected value %+v. Sphere Radius is %f", result, expected, sphere.Radius)
	}
}
