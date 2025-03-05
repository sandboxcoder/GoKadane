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

func TestMax(t *testing.T) {
	center := Vector3{1, 1, 1}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	max := b1.GetMax()
	expected := Vector3{6, 6, 6}
	if max != expected {
		t.Errorf("GetMax() != expected, returned %+v instead of expected value %f", max, expected)
	}
}

func TestMin(t *testing.T) {
	center := Vector3{1, 1, 1}
	b1 := Box{Center: center, Extent: Vector3{X: 5.0, Y: 5.0, Z: 5.0}}
	result := b1.GetMin()
	expected := Vector3{-4, -4, -4}
	if result != expected {
		t.Errorf("result != expected, returned %+v instead of expected value %f", result, expected)
	}
}

func TestExtent(t *testing.T) {
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
