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

// Verify collision primitive with an internal sphere shape can properly be restored.
func TestCollisionPrimitive_RestoreSphere(t *testing.T) {
	s := Sphere{Vector3{0, 0, 0}, 5}
	a := s.GetPrimitive()
	s1 := a.GetSphere()
	if s1 != s {
		t.Errorf("result != expected")
	}
}

func TestCollisionPrimitive_SphereRadiusCheck(t *testing.T) {
	s := Sphere{Vector3{0, 0, 0}, 5}
	a := s.GetPrimitive()
	s1 := a.GetSphere()

	box := Box{Vector3{0, 0, 0}, Vector3{5, 5, 5}}
	b := box.GetPrimitive()
	// Bounding box will return an enlarged sphere
	s2 := b.GetSphere()
	result := s2.Radius == s1.Radius
	expect := false
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}

func TestCollisionPrimitive_SphereVsBox(t *testing.T) {
	s := Sphere{Vector3{0, 0, 0}, 5}
	a := s.GetPrimitive()

	box := Box{Vector3{0, 0, 0}, Vector3{5, 5, 5}}
	b := box.GetPrimitive()

	result := a.IsColliding(b)
	expect := true
	if result != expect {
		t.Errorf("result != expected, returned %+v", result)
	}
}
