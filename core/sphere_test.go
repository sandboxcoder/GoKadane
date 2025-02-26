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

func TestIsCollidingWithPlane(t *testing.T) {
	testSphere := Sphere{Vector3{0, 0, 0}, 20}
	normal := Vector3{0.0, 1.0, 0.0}
	p := Plane{Normal: normal, Dist: 0}
	result := testSphere.isColliding(p)
	if result != true {
		t.Errorf("sphere.isColliding() != expected, expected it to be on the plane")
	}
}

func TestIsCollidingWithPlane_Separated(t *testing.T) {
	testSphere := Sphere{Vector3{0, 90, 0}, 5}
	normal := Vector3{0.0, 1.0, 0.0}
	p := Plane{Normal: normal, Dist: 0}
	result := testSphere.isColliding(p)
	expected := false
	if result != expected {
		t.Errorf("sphere.isColliding() != expected, expected it to not be on the plane")
	}
}
