package core

import (
	"testing"
)

func TestSphereCollision(t *testing.T) {
	s1 := Sphere{Vector3{0, 0, 0}, 20}
	detectedCollision := checkCollision(s1, s1)
	if detectedCollision != true {
		t.Errorf("checkCollision() != expected, returned %+v instead of true", detectedCollision)
	}
}

func TestSphereCollision_Seperated(t *testing.T) {
	s1 := Sphere{Vector3{0, 0, 0}, 20}
	s2 := Sphere{Vector3{100, 0, 0}, 20}
	detectedCollision := checkCollision(s1, s2)
	expected := false
	if detectedCollision != expected {
		t.Errorf("checkCollision() != expected, returned %+v instead of %+v", detectedCollision, expected)
	}
}
