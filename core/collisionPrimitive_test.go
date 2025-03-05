package core

import (
	"testing"
)

func TestCollisionPrimitiveOverlaps_SameObject(t *testing.T) {
	s := Sphere{Vector3{0, 0, 0}, 5}
	a := s.GetPrimitive()
	b := s.GetPrimitive()
	result := a.IsColliding(b)
	expect := true
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}

func TestCollisionPrimitiveOverlaps_NotIntersecting(t *testing.T) {
	s1 := Sphere{Vector3{0, 0, 0}, 5}
	a := s1.GetPrimitive()

	s2 := Sphere{Vector3{90, 0, 0}, 5}
	b := s2.GetPrimitive()
	result := a.IsColliding(b)
	expect := false
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}
