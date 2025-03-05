package core

import (
	"testing"
)

func TestOverlaps(t *testing.T) {
	center := Vector3{0, 0, 0}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	detectedCollision := b1.Overlaps(b1)
	if detectedCollision != true {
		t.Errorf("overlaps() != expected, returned %+v instead of true", detectedCollision)
	}
}

func TestExtent(t *testing.T) {
	center := Vector3{0, 0, 0}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	result := b1.Size()
	max := b1.GetMax()
	min := b1.GetMin()
	expected := (max.Sub(min)).Mag()
	if result != expected {
		t.Errorf("Size() != expected, returned %+v instead of expected value %f", result, expected)
	}
}
